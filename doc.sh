#!/bin/zsh

gh release list --json isDraft,tagName | jq  ".[] | select(.isDraft == false) | .tagName " -r > versions.MD