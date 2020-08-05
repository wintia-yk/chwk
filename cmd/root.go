/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
        "net/http"
        "net/url"
        "bytes"
        "io/ioutil"
)

var token = ""
var roomId = ""
var message = ""
var messageFile = ""

var rootCmd = &cobra.Command{
	Use:   "chwk",
	Short: "Short comment",
	Long: "Long comment",
	Run: func(cmd *cobra.Command, args []string) {
            if len(token) == 0 {
                fmt.Println("Chatwork access token not found.")
                os.Exit(1)
            }
            if len(roomId) == 0 {
                fmt.Println("Chatwork room id not found.")
                os.Exit(1)
            }
            if len(message) == 0 && len(messageFile) == 0 {
                fmt.Println("Message or message file not found.")
                os.Exit(1)
            }

            fmt.Println("Access token: " + token)
            fmt.Println("Room id: " + roomId)
            fmt.Println("Message: " + message)
            fmt.Println("Message file: " + messageFile)

            endpoint := "https://api.chatwork.com/v2"
            requrl := endpoint + "/rooms/" + roomId + "/messages"
            param := ""
            if len(messageFile) == 0 {
                param = "body=" + url.QueryEscape(message)
            } else {
                bytes, err := ioutil.ReadFile(messageFile)
                if err != nil {
                    panic(err)
                }
                param = "body=" + url.QueryEscape(string(bytes))
            }
            request, error := http.NewRequest("POST", requrl, bytes.NewBufferString(param))
            if error != nil {
		fmt.Println(error)
                os.Exit(1)
	    }
            request.Header.Add("X-ChatWorkToken", token)
            request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
            response, error := http.DefaultClient.Do(request)
	    if error != nil {
		fmt.Println(error)
                os.Exit(1)
	    }
            fmt.Println(response)
        },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
        rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Chatwork access token.")
        rootCmd.PersistentFlags().StringVarP(&roomId, "room-id", "r", "", "Chatwork room id.")
        rootCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "Message.")
        rootCmd.PersistentFlags().StringVarP(&messageFile, "message-file", "f", "", "Message file path.")
}

