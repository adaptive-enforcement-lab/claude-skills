# Enable multiple concurrent ephemeral runners
systemctl enable github-runner-ephemeral@{1..5}.service
systemctl start github-runner-ephemeral@{1..5}.service