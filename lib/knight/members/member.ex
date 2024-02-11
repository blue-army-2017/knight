defmodule Knight.Members.Member do
  use Knight.Schema
  import Ecto.Changeset

  schema "members" do
    field :active, :boolean, default: false
    field :firstName, :string
    field :lastName, :string

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(member, attrs) do
    member
    |> cast(attrs, [:firstName, :lastName, :active])
    |> validate_required([:firstName, :lastName, :active])
  end
end
