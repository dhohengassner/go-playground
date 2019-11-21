package kubernetes

import (
	"bufio"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
)

/*
ResourceParser struct to resolve kubernetes resources from files
*/
type ResourceParser struct {
}

/*
NewResourceParser constructor for config *ResourceParser
*/
func NewResourceParser() *ResourceParser {
	return &ResourceParser{}
}

/*
ParseYAML giving back the interface provided by kubernetes client-go
*/
func (rp *ResourceParser) ParseYAML(path string) (obj []interface{}, err error) {

	f, err := os.Open(filepath.Clean(path))
	defer f.Close()
	if err != nil {
		return
	}
	reader := bufio.NewReader(f)
	s := k8syaml.NewYAMLReader(reader)
	for {
		b, ok := s.Read()
		if ok == io.EOF {
			break
		} else if ok != nil {
			return obj, ok
		}

		sobj, _, ok := scheme.Codecs.UniversalDeserializer().Decode(b, nil, nil)
		if ok != nil {
			log.Errorf("not able to parse the part of a YAML: %s", string(b))
		} else {
			obj = append(obj, sobj)
		}
	}

	return
}

// GetImages parses all images from a path (file or dir) and return a list of used images
func (rp *ResourceParser) GetImages(path string) (images []string, err error) {

	objArr, err := rp.ParseYAML(path)
	if err != nil {
		return
	}

	for _, obj := range objArr {
		switch res := obj.(type) {
		case *appsv1.Deployment:
			for _, c := range res.Spec.Template.Spec.Containers {
				images = append(images, c.Image)
			}
			for _, c := range res.Spec.Template.Spec.InitContainers {
				images = append(images, c.Image)
			}
		case *appsv1beta1.Deployment:
			for _, c := range res.Spec.Template.Spec.Containers {
				images = append(images, c.Image)
			}
			for _, c := range res.Spec.Template.Spec.InitContainers {
				images = append(images, c.Image)
			}
		case *appsv1beta2.Deployment:
			for _, c := range res.Spec.Template.Spec.Containers {
				images = append(images, c.Image)
			}
			for _, c := range res.Spec.Template.Spec.InitContainers {
				images = append(images, c.Image)
			}
		case *extensionsv1beta1.Deployment:
			for _, c := range res.Spec.Template.Spec.Containers {
				images = append(images, c.Image)
			}
			for _, c := range res.Spec.Template.Spec.InitContainers {
				images = append(images, c.Image)
			}
		case *appsv1.StatefulSet:
			for _, c := range res.Spec.Template.Spec.Containers {
				images = append(images, c.Image)
			}
			for _, c := range res.Spec.Template.Spec.InitContainers {
				images = append(images, c.Image)
			}
		case *appsv1beta1.StatefulSet:
			for _, c := range res.Spec.Template.Spec.Containers {
				images = append(images, c.Image)
			}
			for _, c := range res.Spec.Template.Spec.InitContainers {
				images = append(images, c.Image)
			}
		case *appsv1beta2.StatefulSet:
			for _, c := range res.Spec.Template.Spec.Containers {
				images = append(images, c.Image)
			}
			for _, c := range res.Spec.Template.Spec.InitContainers {
				images = append(images, c.Image)
			}
		case *corev1.Pod:
			for _, c := range res.Spec.Containers {
				images = append(images, c.Image)
			}
			for _, c := range res.Spec.InitContainers {
				images = append(images, c.Image)
			}
		default:
			log.Warnf("type for Deployment file currently not supported for type %T", res)
		}
	}
	return
}
