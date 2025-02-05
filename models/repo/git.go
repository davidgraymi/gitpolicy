// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo

import (
	"context"

	"code.gitea.io/gitea/models/db"
)

// MergeStyle represents the approach to merge commits into base branch.
type MergeStyle string

const (
	// MergeStyleMerge create merge commit
	MergeStyleMerge MergeStyle = "merge"
	// MergeStyleRebase rebase before merging, and fast-forward
	MergeStyleRebase MergeStyle = "rebase"
	// MergeStyleRebaseMerge rebase before merging with merge commit (--no-ff)
	MergeStyleRebaseMerge MergeStyle = "rebase-merge"
	// MergeStyleSquash squash commits into single commit before merging
	MergeStyleSquash MergeStyle = "squash"
	// MergeStyleFastForwardOnly fast-forward merge if possible, otherwise fail
	MergeStyleFastForwardOnly MergeStyle = "fast-forward-only"
	// MergeStyleManuallyMerged pr has been merged manually, just mark it as merged directly
	MergeStyleManuallyMerged MergeStyle = "manually-merged"
	// MergeStyleRebaseUpdate not a merge style, used to update pull head by rebase
	MergeStyleRebaseUpdate MergeStyle = "rebase-update-only"
)

// MergeStrategy represents the strategy for merging.
type MergeStrategy string

const (
	MergeStrategyDefault   MergeStrategy = ""
	MergeStrategyOrt       MergeStrategy = "ort"
	MergeStrategyRecursive MergeStrategy = "recursive"
	MergeStrategyResolve   MergeStrategy = "resolve"
	MergeStrategyOctopus   MergeStrategy = "octopus"
	MergeStrategyOurs      MergeStrategy = "ours"
	MergeStrategySubtree   MergeStrategy = "subtree"
)

// MergeStrategyOption represents the options for merging strategies.
type MergeStrategyOption string

const (
	MergeStrategyOptionNone              MergeStrategyOption = ""
	MergeStrategyOptionOurs              MergeStrategyOption = "ours"
	MergeStrategyOptionTheirs            MergeStrategyOption = "theirs"
	MergeStrategyOptionRenormalize       MergeStrategyOption = "renormalize"
	MergeStrategyOptionNoRenormalize     MergeStrategyOption = "no-renormalize"
	MergeStrategyOptionIgnoreSpaceChange MergeStrategyOption = "ignore-space-change"
	MergeStrategyOptionIgnoreAllSpace    MergeStrategyOption = "ignore-all-space"
	MergeStrategyOptionIgnoreSpaceAtEol  MergeStrategyOption = "ignore-space-at-eol"
	MergeStrategyOptionIgnoreCrAtEol     MergeStrategyOption = "ignore-cr-at-eol"
)

// UpdateDefaultBranch updates the default branch
func UpdateDefaultBranch(ctx context.Context, repo *Repository) error {
	_, err := db.GetEngine(ctx).ID(repo.ID).Cols("default_branch").Update(repo)
	return err
}
