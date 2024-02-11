defmodule Knight.MembersFixtures do
  @moduledoc """
  This module defines test helpers for creating
  entities via the `Knight.Members` context.
  """

  @doc """
  Generate a member.
  """
  def member_fixture(attrs \\ %{}) do
    {:ok, member} =
      attrs
      |> Enum.into(%{
        active: true,
        firstName: "some firstName",
        lastName: "some lastName"
      })
      |> Knight.Members.create_member()

    member
  end
end
