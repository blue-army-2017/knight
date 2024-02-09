defmodule Knight.Repo do
  use Ecto.Repo,
    otp_app: :knight,
    adapter: Ecto.Adapters.SQLite3
end
