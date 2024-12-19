import { useState } from "react";
import Button from "../Button";
import Input from "../Input";
import Switch from "../Switch";
import { useCreateUser, useUpdateUser, useDeleteUser } from "../../api/user";
import { useQueryClient } from "@tanstack/react-query";
import { User } from "../../types/user";

type UserFormData = {
  name: string;
  age: string;
  active: boolean;
};

interface UserFormProps {
  selectedUserId: string | null;
  selectedUser?: User;
}

export default function UserForm({ selectedUserId, selectedUser }: UserFormProps) {
  const [formData, setFormData] = useState<UserFormData>({
    name: selectedUser?.Name ?? "",
    age: selectedUser?.Age.toString() ?? "",
    active: selectedUser?.Active ?? true,
  });

  const { mutateAsync: createUser } = useCreateUser();
  const { mutateAsync: updateUser } = useUpdateUser();
  const { mutateAsync: deleteUser } = useDeleteUser();

  const queryClient = useQueryClient();

  const handleSubmit = async () => {
    if (selectedUserId) {
      await updateUser({
        id: selectedUserId,
        name: formData.name,
        age: formData.age,
        active: formData.active,
      });
    } else {
      await createUser({
        name: formData.name,
        age: formData.age,
        active: formData.active,
      });
    }

    queryClient.invalidateQueries({ queryKey: ["users"] });

    setFormData({
      name: "",
      age: "",
      active: true,
    });
  };

  const handleDelete = async () => {
    if (selectedUserId) {
      await deleteUser(selectedUserId);

      queryClient.invalidateQueries({ queryKey: ["users"] });

      setFormData({
        name: "",
        age: "",
        active: true,
      });
    }
  };

  const isFormValid = formData.name && formData.age;

  const isUpdateButtonDisabled = !selectedUserId || !isFormValid;

  const isDeleteButtonDisabled = !selectedUserId;

  return (
    <div className="flex flex-col gap-4 p-2">
      <div className="flex flex-row items-center gap-2">
        <Input
          placeholder="Nome"
          value={formData.name}
          onChange={(e) => setFormData({ ...formData, name: e.target.value })}
        />
        <Input
          placeholder="Idade"
          type="number"
          value={formData.age}
          onChange={(e) => setFormData({ ...formData, age: e.target.value })}
        />
        <Switch
          checked={formData.active}
          onCheckedChange={(e) => setFormData({ ...formData, active: e })}
        />
      </div>

      <div className="flex justify-between">
        <Button
          title={"Registrar usuário"}
          mode="primary"
          onClick={handleSubmit}
          className="max-w-64"
          disabled={!isFormValid || !!selectedUserId}
        />

        <Button
          title={"Atualizar usuário"}
          mode="primary"
          onClick={handleSubmit}
          className="max-w-64"
          disabled={isUpdateButtonDisabled}
        />

        <Button
          title="Remover usuário"
          mode="primary"
          onClick={handleDelete}
          className="!bg-red-600 !border-red-600 max-w-64"
          disabled={isDeleteButtonDisabled}
        />
      </div>
    </div>
  );
}
