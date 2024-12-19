import clsx from "clsx";
import {
  InputHTMLAttributes,
  forwardRef,
} from "react";

export interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  error?: string;
}

const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ onChange = () => {}, onBlur = () => {}, error, value, ...props }, ref) => {
    return (
      <>
        <input
          ref={ref}
          {...props}
          className={clsx(
            "w-full flex h-8 items-center justify-between rounded-md border px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus:outline-none disabled:cursor-not-allowed disabled:bg-[#f3f3f3] transition-colors group",
            {
              "border-red-500": error,
              "border-zinc-300 focus:border-dark_blue_mk": !error,
            },
            props.className,
          )}
          onChange={onChange}
          onBlur={onBlur}
          value={value}
          placeholder={props.placeholder || "Clique para preencher"}
        />
        {error && <span className="text-xs text-red_mk">{error}</span>}
      </>
    );
  },
);

export default Input;
