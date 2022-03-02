/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	versionsLib "github.com/kevinmgranger/crev/pkg/versions"
	"github.com/spf13/cobra"
)

// versioninfoCmd represents the versioninfo command
var versioninfoCmd = &cobra.Command{
	Use:   "versioninfo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run:  vinfo,
}

func printOneVersion(version string) {
	var info versionsLib.ReleaseInfo

	if offline {
		info = getOneOffline(version)
	} else {
		info = getOneOnline(version)
	}
	prettyJsonTo(os.Stdout, info)
}

func getOneOffline(version string) versionsLib.ReleaseInfo {

	if info, ok := versionsLib.KnownVersionInfo[version]; ok {
		return info
	}
	dief("Unknown version %v", version)
	panic("unreachable")
}

func getOneOnline(version string) versionsLib.ReleaseInfo {
	info, err := versionsLib.GetInfoForVersion(version)

	if err != nil {
		die(err)
	}
	return info
}

func printManyVersions(versions []string) {
	infos := make(map[string]versionsLib.ReleaseInfo, len(versions))

	var get func(version string) versionsLib.ReleaseInfo

	if offline {
		get = getOneOffline
	} else {
		get = getOneOnline
	}

	for _, version := range versions {
		if version == "all" {
			die("Cannot use 'all' with multiple versions")
		}

		infos[version] = get(version)
	}

	prettyJsonTo(os.Stdout, infos)
}

func printAll() {
	var versions map[string]versionsLib.ReleaseInfo
	if offline {
		versions = versionsLib.KnownVersionInfo
	} else {
		var err error
		versions, err = versionsLib.GetAllVersionInfo()
		if err != nil {
			die(err)
		}
	}
	prettyJsonTo(os.Stdout, versions)

}

func vinfo(cmd *cobra.Command, args []string) {
	switch len(args) {
	case 1:
		version := args[0]
		if version == "all" {
			printAll()
		} else {
			printOneVersion(version)
		}
	default:
		printManyVersions(args)
	}
}

var offline bool

func init() {
	versioninfoCmd.Flags().BoolVar(&offline, "offline", false, "")
	getCmd.AddCommand(versioninfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versioninfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versioninfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
