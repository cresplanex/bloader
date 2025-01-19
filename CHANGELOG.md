# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v1.0.6] - 2025-01-19

### 🐞 Bug fixes
✔️ Fix depend_on flow (omniarchy <k.hayashi@cresplanex.com>) [2025-01-20 04:44:59 +0900]


## [v1.0.5] - 2025-01-19

### 🐞 Bug fixes
✔️ Avoid temporary panic with chan for wait close (omniarchy <k.hayashi@cresplanex.com>) [2025-01-19 23:49:16 +0900]


### 💻 Chores
✔️ Update with_slave example case (omniarchy <k.hayashi@cresplanex.com>) [2025-01-16 00:48:48 +0900]


## [v1.0.4] - 2025-01-15

### 🐞 Bug fixes
✔️ No err on ctx cancel (omniarchy <k.hayashi@cresplanex.com>) [2025-01-15 21:19:03 +0900]

✔️ Fix count limit behavior on mass exec (omniarchy <k.hayashi@cresplanex.com>) [2025-01-15 21:18:45 +0900]


## [v1.0.3] - 2025-01-15

### 🐞 Bug fixes
✔️ Prevent exec + 1 count (omniarchy <k.hayashi@cresplanex.com>) [2025-01-15 19:36:22 +0900]


### 🔔 Others
✔️ Fix: Fix dont need print (omniarchy <k.hayashi@cresplanex.com>) [2025-01-15 04:38:23 +0900]


## [v1.0.2] - 2025-01-14

### 🚀 Features
✔️ Add slave setup with terraform #32 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 20:44:35 +0900]

✔️ Merge pull request #31 from cresplanex/feat/feature-add-program-info-to-log_#30 (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-14 19:55:40 +0900]

✔️ Add code Info to Logger #30 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 19:50:22 +0900]

✔️ Delete on property represents FunctionName #30 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 19:49:53 +0900]

✔️ Merge pull request #29 from cresplanex/feat/feature-define-task-for-automation_#28 (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-14 19:37:49 +0900]

✔️ Merge pull request #27 from cresplanex/feat/bug-panic-by-type-of-atomicerr_#26 (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-14 18:10:01 +0900]

✔️ Define sync common error #26 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 17:49:59 +0900]


### 🐞 Bug fixes
✔️ fix; Update Runner definition #32 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 23:04:15 +0900]

✔️ Use atomic Pointer from atomic Value #26 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 17:50:09 +0900]


### 📃 Documentation
✔️ Merge pull request #33 from cresplanex/docs/docs-creating-an-example_#32 (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-14 23:06:53 +0900]

✔️ Create Example guide #32 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 23:04:32 +0900]

✔️ Add with_slave bloader.yaml #32 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 20:36:40 +0900]

✔️ Add terraform cmd on asdf #32 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 20:34:05 +0900]

✔️ Add task automation docs #28 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 19:26:28 +0900]


### 💻 Chores
✔️ Create question.md (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-14 20:01:50 +0900]

✔️ Create security_issue.md (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-14 20:01:12 +0900]

✔️ Create documentation_request.md (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-14 20:00:09 +0900]

✔️ Create maintenance_talk.md (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-14 19:59:23 +0900]

✔️ Changed formatting method to not include automatically generated go files in proto #28 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 19:26:09 +0900]

✔️ Define vscode task for task flow #28 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 19:24:55 +0900]

✔️ Create Makefile for task flow #28 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 19:24:37 +0900]


### 🔔 Others
✔️ Add runner with_slave example #32 (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 20:37:11 +0900]


## [v1.0.1] - 2025-01-14

### 📃 Documentation
✔️ Update README.md (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-14 03:23:07 +0900]


### 💻 Chores
✔️ Fix go releaser format (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 11:52:27 +0900]


## [v1.0.0] - 2025-01-13

### 🚀 Features
✔️ Go mod tidy (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 23:43:55 +0900]

✔️ Change Dep Lib for proto from remote to local (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 23:43:42 +0900]

✔️ Add Version prop to config (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 23:42:43 +0900]

✔️ Gen proto code (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 23:41:46 +0900]

✔️ Add go_package to proto (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 23:41:17 +0900]

✔️ Define buf gen yaml (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 23:40:53 +0900]

✔️ Add API Key auth msg (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 22:41:51 +0900]

✔️ Add basic auth msg (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 22:34:29 +0900]

✔️ Add main feature (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 19:31:53 +0900]


### 🐞 Bug fixes
✔️ Exec goimports (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 23:50:22 +0900]

✔️ Fix proto go_package (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 23:43:05 +0900]

✔️ Trigger CI flow on only main push (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 22:29:28 +0900]

✔️ Update buf mod (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 21:57:29 +0900]


### 📃 Documentation
✔️ Update doc related buf proto (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 00:04:40 +0900]

✔️ Add Bud doc (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 23:50:34 +0900]

✔️ Update logo link (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-13 21:49:04 +0900]

✔️ Update README.md (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-13 21:22:48 +0900]

✔️ Update README.md logo size (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-13 21:21:24 +0900]

✔️ Update README.md margin (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-13 21:20:21 +0900]

✔️ Add margin under the logo (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 21:13:50 +0900]

✔️ Update logo (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 21:11:01 +0900]

✔️ Upload logo (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 20:51:15 +0900]

✔️ Update repository owner on installation.md (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 19:39:41 +0900]

✔️ Update mailaddress, name (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 19:34:55 +0900]


### 🕶️ Styles
✔️ Formatting buf (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 23:53:27 +0900]

✔️ Exec buf format (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 22:46:04 +0900]

✔️ Fix format (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 22:38:18 +0900]

✔️ Formatting root.go (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 20:01:22 +0900]


### 🧪 Tests
✔️ Test 2 CI flow (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 22:31:37 +0900]

✔️ Test CI flow (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 22:31:20 +0900]


### 💻 Chores
✔️ Update `update changelog` flow (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 00:45:27 +0900]

✔️ Add buf CLI trigger on tag push (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 00:31:01 +0900]

✔️ Fix update changelog flow on first release (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 00:29:01 +0900]

✔️ Add trigger CI flow on develop push (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 22:30:28 +0900]

✔️ Add edit conf (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 22:28:20 +0900]

✔️ chore; Change CI trigger (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 20:51:29 +0900]

✔️ Add dependabot (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 19:58:40 +0900]

✔️ Add update changelog action (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 19:58:26 +0900]

✔️ Add release flow (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 19:58:05 +0900]

✔️ Add CI flow (omniarchy <k.hayashi@cresplanex.com>) [2025-01-13 19:57:38 +0900]


### 🔔 Others
✔️ Merge remote-tracking branch 'refs/remotes/origin/main' (omniarchy <k.hayashi@cresplanex.com>) [2025-01-14 00:46:05 +0900]

✔️ Merge pull request #11 from cresplanex/develop (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-14 00:33:16 +0900]

✔️ Merge pull request #10 from cresplanex/develop (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-14 00:08:56 +0900]

✔️ Merge pull request #9 from cresplanex/develop (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-13 22:39:44 +0900]

✔️ Merge pull request #7 from cresplanex/develop (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-13 22:03:33 +0900]

✔️ Merge pull request #1 from cresplanex/develop (HAYASHI KENTA <83863286+ablankz@users.noreply.github.com>) [2025-01-13 21:39:55 +0900]
