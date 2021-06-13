package main

import (
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func init() {
	initDOConfig() // Initialize digitalocean configs using env vars
}

// init digital ocean config based on env
func initDOConfig() {
	// os.Setenv("DIGITALOCEAN_API_URL")
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cluster, err := digitalocean.NewKubernetesCluster(
			ctx,
			"polkadot-k8s",
			&digitalocean.KubernetesClusterArgs{
				NodePool: &digitalocean.KubernetesClusterNodePoolArgs{
					Name:      pulumi.String("polkadot-k8s-np"),
					Size:      pulumi.String("s-2vcpu-2gb"),
					NodeCount: pulumi.Int(3)},
				Region:  pulumi.String("AMS3"),
				Version: pulumi.String("1.20.7-do.0"),
			})
		if err != nil {
			return err
		}
		ctx.Export("configs", cluster.KubeConfigs)
		return nil
	})
}
