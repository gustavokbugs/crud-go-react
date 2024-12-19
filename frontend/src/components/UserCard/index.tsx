import { User } from "../../types/user";

interface UserCardProps {
  user: User;
  isSelected: boolean;
  onSelect: (id: string) => void;  // Modificação: agora onSelect recebe o id
}

export default function UserCard({ user, isSelected, onSelect }: UserCardProps) {
  return (
    <div
      className={`flex flex-col gap-2 bg-gray-200 m-4 p-4 rounded-md min-w-64 ${isSelected ? 'border-4 border-blue-500' : ''}`}
      onClick={() => onSelect(user.ID)}  // Passa o id do usuário ao clicar
    >
      <span className="font-bold text-lg">{user.Name}</span>
      <hr className="h-[1.5px] bg-black w-full" />
      <div className="flex flex-row items-center gap-2">
        <span className="font-bold">Idade:</span>
        <span className="text-gray-700">{user.Age}</span>
      </div>
      <div className="flex flex-row items-center gap-2">
        <span className="font-bold">Ativo:</span>
        <span className="text-gray-700">{user.Active ? "Sim" : "Não"}</span>
      </div>
      <div className="flex flex-row items-center gap-2">
        <span className="font-bold">ID:</span>
        <span className="text-gray-700">{user.ID}</span>
      </div>
    </div>
  );
}
