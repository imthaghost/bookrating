// Copyright 2022 imthaghost. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
)

var (
	// Root cmd
	rootCmd = &cobra.Command{
		Use:   "bookrating <file path>",
		Short: "Upload reviews to Notion with ease!",
		Long:  `Has your book club been recording their favorite data for 10 years and have no way of reading it? Look no further! Bookrating will ingest a given CSV of book reviews and present it to you in a beautiful format using Notion.`,
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// Print the usage if no args are passed in :)
			if len(args) < 1 {
				if err := cmd.Usage(); err != nil {
					log.Fatal(err)
				}

				return
			}

			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
			defer stop()
			// Otherwise.. generate reviews!
			if err := generateReview(ctx, args); err != nil {
				log.Fatalf("%+v", err)
			}
		},
	}
)

// Execute the clone command
func Execute() {
	// Persistent Flags
	// none

	// Execute the command :)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}