defmodule Knight.Repo.Migrations.CreateMembers do
  use Ecto.Migration

  def change do
    create table(:members, primary_key: false) do
      add :id, :string, primary_key: true
      add :firstName, :string, null: false
      add :lastName, :string, null: false
      add :active, :boolean, default: false, null: false

      timestamps(type: :utc_datetime)
    end
  end
end
