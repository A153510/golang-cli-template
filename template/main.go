/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

{%- if values.isAzdo %}
import "dev.azure.com/${{ values.repoOrg }}/${{ values.repoOwner }}/_git/${{ values.repoName }}/cmd"
{%- else %}
import "github.com/${{ values.repoOwner }}/${{ values.repoName }}/cmd"
{%- endif %}

func main() {
	cmd.Execute()
}
