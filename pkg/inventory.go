package cfg

import (
        "encoding/json"
	"io/ioutil"
	"net/http"
        "os"
	"crypto/tls"
        "github.com/tidwall/gjson"
        b64 "encoding/base64"
        "github.com/ghodss/yaml"
)

func cfg_kubernetes_list(cfgurl string, token string) ([]gjson.Result, error) {

    cfgurl_clusters := cfgurl + "/kubernetes/clusters/cluster/"

    data, err := make_get_request(cfgurl_clusters, token)
    if err != nil {
        return nil, err
    }
    clusters_list := gjson.Get(string(data), "clusters")
    clusters_lists := clusters_list.Array()

    return clusters_lists, nil
}

func cfg_kubernetes_create_dict(clusters_lists []gjson.Result, cfgurl string, token string, clientid string, clientsecret string, clientca string) ([]byte, error) {
    json_clusters := []map[string]interface{}{}
    cfgurl_clusters := cfgurl + "/kubernetes/clusters/cluster/"    

    for _, cluster := range clusters_lists {
        cfgurl_kubeconfig := cfgurl_clusters + cluster.String() + "/kubeconfig/"
        data_kubeconfig, errapi := make_get_request(cfgurl_kubeconfig, token)
        if errapi != nil {
            return nil, errapi
        }
        kubeconfig_b64 := gjson.Get(string(data_kubeconfig), "kubeconfig")
        kubeconfig, _ := b64.StdEncoding.DecodeString(kubeconfig_b64.String())
        json_kube, err := yaml.YAMLToJSON(kubeconfig)
        if err != nil {
            return nil, err
        }
        cluster_url := gjson.Get(string(json_kube), "clusters.0.cluster.server")
        certificate := gjson.Get(string(json_kube), "clusters.0.cluster.certificate-authority-data")
        certificate_decrypt, _ := b64.StdEncoding.DecodeString(certificate.String())

        errca := os.WriteFile("/tmp/ca-"+cluster.String(), certificate_decrypt, 0600)
        if errca != nil {
            return nil, errca
        }
        json_cluster := map[string]interface{}{}
        json_cluster["name"] = cluster.String()
        json_cluster["apiServer"] = map[string]interface{}{}
        json_cluster["apiServer"].(map[string]interface{})["url"] = cluster_url.String()
        json_cluster["apiServer"].(map[string]interface{})["caFile"] = "/tmp/ca-"+cluster.String()
        json_cluster["oauth"] = map[string]interface{}{}
        json_cluster["oauth"].(map[string]interface{})["clientID"] = clientid
        json_cluster["oauth"].(map[string]interface{})["clientSecret"] = clientsecret
        json_cluster["oauth"].(map[string]interface{})["caFile"] = clientca
        
        json_clusters = append(json_clusters, json_cluster)
    }
    b, err := json.Marshal(json_clusters)
    if err != nil {
        return nil, err
    }

    return b, nil

}

func make_get_request(cfgurl string, token string) ([]byte, error) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", cfgurl, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Authorization", "Token " + token)

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    client = &http.Client{Transport: tr}
    response, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    data, _ := ioutil.ReadAll(response.Body)
    return data, nil
}

func Get_inventory(url string, token string, clientid string, clientsecret string, clientca string) ([]byte, error) {


  clusters, err := cfg_kubernetes_list(url, token)
  if err != nil {
      return nil, err 
  }
  result, err := cfg_kubernetes_create_dict(clusters, url, token, clientid, clientsecret, clientca)
  if err != nil {
      return nil, err
  }

  return result, nil

}

