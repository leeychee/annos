// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/leeychee/annos"
	"github.com/leeychee/annos/detrac"
	"github.com/spf13/cobra"
)

var (
	srcFmt string
	dstFmt string
	output string
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert the given format to the expected format.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if srcFmt == "detrac" && dstFmt == "ssd" {
			fmt.Printf("Convert %s -> %s\n", srcFmt, dstFmt)
			if len(args) < 1 {
				fmt.Println("need detrac annotation filepath")
				return

			}
			src := args[0]
			dst := output
			if len(args) > 1 {
				dst = args[1]
			}
			convertDetrac2Ssd(src, dst)
		} else {
			fmt.Printf("Unsupported format: %s -> %s\n", srcFmt, dstFmt)
		}
	},
}

func init() {
	RootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&srcFmt, "srcfmt", "s", "", "the source format.")
	convertCmd.Flags().StringVarP(&dstFmt, "dstfmt", "d", "", "the destination format.")
	convertCmd.Flags().StringVarP(&output, "output", "o", ".", "output folder")
}

func convertDetrac2Ssd(src, dst string) {
	fmt.Printf("Read sequence from %s\n", src)
	seq, err := detrac.ReadFromFile(src)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Get %d DETRAC frames.\n", len(seq.Frames))
	ss, err := annos.Detrac2Ssd(seq)
	if err != nil {
		fmt.Println(err)
	}
	// If not exist, created.
	folder := filepath.Join(dst, seq.Name)
	_, err = os.Stat(folder)
	if err != nil {
		fmt.Println(err)
		os.Mkdir(folder, 0777)
	}
	var count int
	for _, s := range ss {
		b, err := xml.MarshalIndent(s, "", "\t")
		if err != nil {
			fmt.Println(err)
		}
		// from
		buf := bytes.NewBufferString(xml.Header)
		buf.Write(b)
		// to
		filename := filepath.Join(folder, s.Filename[:len(s.Filename)-3]+"xml")
		outf, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		io.Copy(outf, buf)
		err = outf.Close()
		if err != nil {
			fmt.Printf("fail to write file: %s, %s\n", filename, err)
		}
		count++
	}
	fmt.Printf("Save %d ssd annotations.\n", count)
}
