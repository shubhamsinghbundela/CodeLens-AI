import React from "react";
import { cn } from "@/lib/utils";

type Props = {
  id?: string;
  className?: string;
  children?: React.ReactNode;
};

const Container = React.forwardRef<HTMLElement, Props>(({ id, className, children }, ref) => {
  return (
    <section ref={ref} id={id} className={cn("mx-auto w-full max-w-372 px-4 md:px-12", className)}>
      {children}
    </section>
  );
});

Container.displayName = "Container";
export default Container;
