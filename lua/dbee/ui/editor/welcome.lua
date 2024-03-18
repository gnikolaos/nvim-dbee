local M = {}

function M.banner()
  return {
    "[ Enter insert mode to clear ]",
    "",
    "",
    "Welcome to",
    "",
    " ██████████   ███████████",
    "░░███░░░░███ ░░███░░░░░███",
    " ░███   ░░███ ░███    ░███  ██████   ██████",
    " ░███    ░███ ░██████████  ███░░███ ███░░███",
    " ░███    ░███ ░███░░░░░███░███████ ░███████",
    " ░███    ███  ░███    ░███░███░░░  ░███░░░",
    " ██████████   ███████████ ░░██████ ░░██████",
    "░░░░░░░░░░   ░░░░░░░░░░░   ░░░░░░   ░░░░░░",
    "",
    "",
    'Type ":h dbee.txt" to learn more about the plugin.',
    "",
    'Report issues to: "github.com/kndndrj/nvim-dbee/issues".',
  }
end

return M
