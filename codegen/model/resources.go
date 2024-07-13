package model

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	PersistentVolumeType = metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
	}
	PersistentVolumeClaimType = metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
	}
)
