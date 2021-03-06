// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DestinationRuleLister helps list DestinationRules.
type DestinationRuleLister interface {
	// List lists all DestinationRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha3.DestinationRule, err error)
	// DestinationRules returns an object that can list and get DestinationRules.
	DestinationRules(namespace string) DestinationRuleNamespaceLister
	DestinationRuleListerExpansion
}

// destinationRuleLister implements the DestinationRuleLister interface.
type destinationRuleLister struct {
	indexer cache.Indexer
}

// NewDestinationRuleLister returns a new DestinationRuleLister.
func NewDestinationRuleLister(indexer cache.Indexer) DestinationRuleLister {
	return &destinationRuleLister{indexer: indexer}
}

// List lists all DestinationRules in the indexer.
func (s *destinationRuleLister) List(selector labels.Selector) (ret []*v1alpha3.DestinationRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.DestinationRule))
	})
	return ret, err
}

// DestinationRules returns an object that can list and get DestinationRules.
func (s *destinationRuleLister) DestinationRules(namespace string) DestinationRuleNamespaceLister {
	return destinationRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DestinationRuleNamespaceLister helps list and get DestinationRules.
type DestinationRuleNamespaceLister interface {
	// List lists all DestinationRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha3.DestinationRule, err error)
	// Get retrieves the DestinationRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha3.DestinationRule, error)
	DestinationRuleNamespaceListerExpansion
}

// destinationRuleNamespaceLister implements the DestinationRuleNamespaceLister
// interface.
type destinationRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DestinationRules in the indexer for a given namespace.
func (s destinationRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.DestinationRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.DestinationRule))
	})
	return ret, err
}

// Get retrieves the DestinationRule from the indexer for a given namespace and name.
func (s destinationRuleNamespaceLister) Get(name string) (*v1alpha3.DestinationRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("destinationrule"), name)
	}
	return obj.(*v1alpha3.DestinationRule), nil
}
