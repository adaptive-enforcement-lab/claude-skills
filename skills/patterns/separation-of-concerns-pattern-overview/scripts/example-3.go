// cmd/deploy/deploy.go - CLI layer
package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "example.com/pkg/deployer"
)

func NewDeployCommand() *cobra.Command {
    var opts deployer.Options

    cmd := &cobra.Command{
        Use:   "deploy",
        Short: "Deploy application to Kubernetes",
        RunE: func(cmd *cobra.Command, args []string) error {
            // Only CLI concerns here: flag parsing, output, exit codes
            d, err := deployer.New(opts)
            if err != nil {
                return fmt.Errorf("initializing deployer: %w", err)
            }

            // Business logic delegated to pkg/
            result, err := d.Deploy(cmd.Context())
            if err != nil {
                return err
            }

            // Output formatting (CLI concern)
            fmt.Printf("Deployed %s to namespace %s\n", result.Name, result.Namespace)
            return nil
        },
    }

    // Flag binding (CLI concern)
    cmd.Flags().StringVar(&opts.Namespace, "namespace", "default", "Kubernetes namespace")
    cmd.Flags().StringVar(&opts.Image, "image", "", "Container image")
    cmd.MarkFlagRequired("image")

    return cmd
}