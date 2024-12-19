import clsx from "clsx";

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  title: string;
  mode: "primary" | "secondary";
  color?: "primary" | "alternate";
  loading?: boolean;
  className?: string;
}

export default function Button({
  title,
  mode,
  color,
  loading,
  className,
  ...rest
}: ButtonProps) {
  return (
    <button
      disabled={rest.disabled || loading}
      className={clsx(
        "min-w-[154px] h-[32px] whitespace-nowrap lg:px-4 flex items-center justify-center font-medium cursor-pointer text-xs md:text-sm py-2 2xl:text-md rounded border hover:opacity-90 hover:transition hover:duration-300 disabled:cursor-not-allowed disabled:!bg-[#858585] disabled:text-white disabled:border-none disabled:border-gray-500",
        {
          "bg-indigo-800 text-white border-indigo-800": mode === "primary",
          "bg-white text-indigo-800 border-indigo-800 border":
            mode === "secondary",
          "!border-blue-400 !bg-blue-400":
            mode === "primary" && color === "alternate",
          "!text-blue-400 !border-blue-400":
            mode === "secondary" && color === "alternate",
          "!bg-[#858585] !text-white !border-none !opacity-[0.6]":
            color === "alternate" && rest.disabled && mode === "secondary",
        },
        className,
      )}
      {...rest}
    >
      {title}
    </button>
  );
}
