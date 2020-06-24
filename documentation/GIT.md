### GIT 

**What is Version Control?**

Version control is a system that records changes to a file or set of files over time so that you can recall specific versions later. So ideally, we can place any file on the computer on version control.
Why is this useful?
A Version Control System (VCS) allows you to revert files back to a previous state, revert the entire project back to a previous state, review changes made over time, see who last modified something or who introduced what and when. 

**What is a Repository ?**

A repository is a collection of source code.

**Git Workflow (just the basic)**

Diagram of a simple Git Workflow

If you consider a file in your Working Directory, it can be in three possible states.
- **It can be staged.** Which means the files with the updated changes are marked to be committed to the local repository but not yet committed.
- **It can be modified.** Which means the files with the updated changes are not yet stored in the local repository.
- **It can be committed.** Which means that the changes you made to your file are safely stored in the local repository.

Commands:

- ```git add``` is a command used to add a file that is in the working directory to the staging area.
- ```git commit``` is a command used to add all files that are staged to the local repository.
- ```git push``` is a command used to add all committed files in the local repository to the remote repository. So in the remote repository, all files and changes will be visible to anyone with access to the remote repository.
- ```git fetch``` is a command used to get files from the remote repository to the local repository but not into the working directory.
- ```git merge``` is a command used to get the files from the local repository into the working directory.
- ```git pull``` is command used to get files from the remote repository directly into the working directory. It is equivalent to a git fetch and a git merge.

**Let’s create our first repo:**
1) Create a GitHub account https://github.com
- Download and install
```
brew install git
git --version
```

2) Set up your local environment following these steps to https://kbroman.org/github_tutorial/pages/first_time.html

3) Create a new repo (Follow this link)
```
 echo "# test" >> README.md
```

4) Initialize git
```
git init
``` 

5) Add files to the Staging Area for commit:
```
git status         # Lists all new or modified files to be committed
git add README.md  # To add a specific file
git add .          # Adds all the files in the local repository and stages them for commit
git status         # Lists all new or modified files to be committed

```

6) Commit the changes you made:
```
git commit -m "first commit"

```

8) Uncommit changes you just made
```
git reset HEAD~1# Remove the most recent commit# Commit again!
git remote add origin git@github.com:yourgithubaccount/test.git
```

9) Add a remote origin and Push:
```
git remote add origin git@github.com:yourgithubaccount/myRepo.git
git push -u origin master
```

10) Cloning a Git Repo:
Locate to the directory you want to clone the repo. Copy the link of the repository you want and enter the following:
```
git clone remote_repository_URL
```

**Collaborating**:
So imagine you and your friend are collaborating on a project. You both are working on the same project files. 
Each time you make some changes and push it into the master repo, your friend has to pull the changes that you pushed 
into the git repo. Meaning to make sure you’re working on the latest version of the git repo each time you start working, 
a git pull command is the way to go

- Mention branch workflow !!

1) Clone the repo that you want to collaborate
Run git pull origin master to make sure those changes are reflected on my local copy of the repo

2) User1 will create a new branch called “my-first-branch”
```
git-training (master) $ git branch my-first-branch

git-training (master) $ git branch
* master
  my-first-branch

git-training (master) $ git checkout my-first-branch
Switched to branch 'my-first-branch'

git-training (my-first-branch) $ git branch
  master
* my-first-branch

```

3) User2 will add some changes to “my-first-branch”

Solving conflicts

4) Pull request & merge










