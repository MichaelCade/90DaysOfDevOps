# Day 53 - Gickup - Keep your repositories safe
[![Watch the video](thumbnails/day53.png)](https://www.youtube.com/watch?v=hKB3XY7oMgo)

 # ONE SENTENCE SUMMARY:
Andy presented Gickup, a tool for backing up Git repositories across various platforms like GitHub, GitLab, Bitbucket, etc., written in Go. He explained its usage, demonstrated its functionality, and showcased its ability to restore deleted repositories.

# MAIN POINTS:
1. Gickup is a tool written by Andy for backing up Git repositories.
2. It supports GitHub, GitLab, Bitbucket, SourceForge, local repositories, and any type of Git repository as long as you can provide access credentials.
3. Automation is simple; once configured, it takes care of everything.
4. It can be run using pre-compiled binaries, Homebrew, Docker, Arch User Repository (AUR), or NYX.
5. Gickup connects to the API of the host service and grabs the repositories you want to back up.
6. You define a source (like GitHub) and specify a destination, which could be local backup, another Git hoster, or a mirror.
7. The configuration is in YAML, where you define the source, destination, structured format for the backup, and whether to create an organization if it doesn't exist.
8. Demonstration included backing up and restoring repositories, mirroring repositories to another Git hoster, and handling accidental repository deletions.
9. Gickup can be kept up-to-date through the presenter's social media accounts or QR code linked to his GitHub account.


# ONE SENTENCE SUMMARY:
Gickup is a tool written in Go, designed to backup and restore Git repositories, allowing for simple automation and secure backups.

# MAIN POINTS:

1. Gickup is a tool that backs up Git repositories, supporting multiple hosting platforms.
2. It can be run using pre-compiled binaries, Homebrew, Docker, or AUR.
3. Gickup connects to the API of the hoster, grabbing all desired repositories and their contents.
4. Configuration is done in YAML, defining sources, destinations, and backup options.
5. Local backups can be created, with an optional structured directory layout.
6. Mirroring to another hosting platform is also possible, allowing for easy repository management.
7. Gickup provides a simple automation solution for backing up Git repositories.

# TAKEAWAYS:

1. Use Gickup to automate the process of backing up your Git repositories.
2. Gickup supports multiple hosting platforms and allows for secure backups.
3. Configure Gickup using YAML files to define sources, destinations, and backup options.
4. Create local backups or mirror repositories to another hosting platform for easy management.
5. Restore deleted repositories by recreating the repository, grabbing the origin, and pushing changes.
6. Use Gickup to keep your Git repositories safe and organized.
7. Consider using Gickup as a part of your DevOps workflow.
