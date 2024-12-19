import { useState } from "react";
import { UserCard, UserForm } from "./components";
import { useGetUsers } from "./api";

export default function App() {
  const { data: userList } = useGetUsers();
  const [selectedUserId, setSelectedUserId] = useState<string | null>(null);

  const handleSelectUser = (id: string) => {
    console.log("Selected User ID:", id);
    setSelectedUserId(id);
  };

  return (
    <div className="flex items-center justify-center h-screen bg-gray-500">
      <div className="w-[600px] h-auto max-h-[900px] overflow-hidden border-solid rounded-lg bg-white">
        <div className="p-4 text-center">
          <h1>DESAFIO - CRUD GO / REACT - COMPARTILHA TECH</h1>
          <UserForm selectedUserId={selectedUserId} />
        </div>

        <div className="overflow-auto max-h-[700px]">
          {userList?.map((user) => (
            <UserCard
              key={user.ID}
              user={user}
              isSelected={user.ID === selectedUserId}
              onSelect={handleSelectUser}
            />
          ))}
        </div>
      </div>
    </div>
  );
}
