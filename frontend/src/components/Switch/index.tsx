import {
  Root,
  Thumb,
  SwitchProps as DefaultSwitchProps
} from "@radix-ui/react-switch"
import clsx from "clsx";
import { forwardRef } from "react";

type SwitchProps = DefaultSwitchProps;

const Switch = forwardRef<HTMLButtonElement, SwitchProps>(({
  ...rest
}, ref) => {
  return (
    <div className="h-8 flex items-center">
      <Root
        className={
          clsx(
            "w-11 h-6 bg-red-600 border border-white rounded-xl relative",
            "data-[state=checked]:bg-green-400"
          )
        }
        ref={ref}
        {...rest}
      >
        <Thumb className={
          clsx(
            "block w-4 h-4 bg-white rounded-lg transition-all duration-100 translate-x-1 will-change-transform",
            "data-[state=checked]:translate-x-[22px] data-[state=checked]:bg-white"
          )
        } />
      </Root>
    </div>
  )
})

export default Switch
