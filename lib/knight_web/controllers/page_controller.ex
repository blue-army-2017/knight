defmodule KnightWeb.PageController do
  use KnightWeb, :controller

  def home(conn, _params) do
    render(conn, :home)
  end
end
