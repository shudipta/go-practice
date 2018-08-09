package docker

import (
	"fmt"
	"k8s.io/kubernetes/pkg/util/parsers"
	"strings"
	"log"
	"net/http"
	reg "github.com/heroku/docker-registry-client/registry"
)

func interfaceTest(str interface{}) {
	fmt.Println("interface test =", str)
}

//func parseImageName(image string) (string, string, string) {
//	registry := "registry-1.docker.io"
//	tag := "latest"
//	var nameParts, tagParts []string
//	var name, port string
//	state := 0
//	start := 0
//	for i, c := range image {
//		if c == ':' || c == '/' || c == '@' || i == len(image)-1 {
//			if i == len(image)-1 {
//				i += 1
//			}
//			part := image[start:i]
//			start = i + 1
//			switch state {
//			case 0:
//				if strings.Contains(part, ".") {
//					registry = part
//					if c == ':' {
//						state = 1
//					} else {
//						state = 2
//					}
//				} else {
//					if c == '/' {
//						start = 0
//						state = 2
//					} else {
//						state = 3
//						name = fmt.Sprintf("library/%s", part)
//					}
//				}
//			case 3:
//				tag = ""
//				tagParts = append(tagParts, part)
//			case 1:
//				state = 2
//				port = part
//			case 2:
//				if c == ':' || c == '@' {
//					state = 3
//				}
//				nameParts = append(nameParts, part)
//			}
//		}
//	}
//
//	if port != "" {
//		registry = fmt.Sprintf("%s:%s", registry, port)
//	}
//
//	if name == "" {
//		name = strings.Join(nameParts, "/")
//	}
//
//	if tag == "" {
//		tag = strings.Join(tagParts, ":")
//	}
//
//	registry = fmt.Sprintf("https://%s", registry)
//
//	return registry, name, tag
//}

func parseImageName(imageName, registryUrl string) (string, string, string, string, error) {
	repo, tag, digest, err := parsers.ParseImageName(imageName)
	if err != nil {
		return "", "", "", "", err
	}
	// the repo part should have registry url as prefix followed by a '/'
	// for example, if image name = "ubuntu" then
	//					repo = "docker.io/library/ubuntu", tag = "latest", digest = ""
	// 				if image name = "k8s.gcr.io/kubernetes-dashboard-amd64:v1.8.1" then
	//					repo = "k8s.gcr.io/kubernetes-dashboard-amd64", tag = "v1.8.1", digest = ""
	// here, for docker registry the api url is "https://registry-1.docker.io"
	// and for other registry the url is "https://k8s.gcr.io"(gcr) or "https://quay.io"(quay)
	parts := strings.Split(repo, "/")
	if registryUrl == "" {
		if parts[0] == "docker.io" {
			registryUrl = "https://registry-1." + parts[0]
		} else {
			registryUrl = "https://" + parts[0]
		}
	}
	repo = strings.Join(parts[1:], "/")

	return registryUrl, repo, tag, digest, err
}

func processImageName(imageName, user, pass string) {
	// TODO: need to check for digest part
	registryUrl := ""
	registryUrl, repo, tag, digest, err := parseImageName(imageName, registryUrl)
	//repo, tag, digest, err := parsers.ParseImageName(imageName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(registryUrl, repo, tag, digest)

	hub := &reg.Registry{
		URL: registryUrl,
		Client: &http.Client{
			Transport: reg.WrapTransport(http.DefaultTransport, registryUrl, user, pass),
		},
		Logf: reg.Quiet,
	}

	manifest, err := hub.ManifestV2(repo, tag)
	if err != nil {
		log.Fatalf("manifest: %v", err)
	}
	canonicalBytes, err := manifest.MarshalJSON()
	fmt.Println(string(canonicalBytes))

	fmt.Println("--------------------------------")
}

func main() {
	fmt.Println("Hello, playground")
	fmt.Println("--------------------------------")
	//processImageName("ubuntu", "", "")
	//processImageName("shudipta/labels", "shudipta", "pi-shudipta")
	//processImageName("kubernetes-dashboard-amd64:v1.8.1")
	processImageName("k8s.gcr.io/kubernetes-dashboard-amd64:v1.8.1", "", "")
	//processImageName("quay.io/coreos/clair:2.0.0")
	//processImageName("quay.io/coreos/clair:2.0.1")

	interfaceTest("aaaaaa")
	interfaceTest(5)
}