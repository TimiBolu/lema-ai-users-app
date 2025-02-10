import { FC, ReactNode } from "react";

type AppLayoutProps = {
  children: ReactNode;
};

const AppLayout: FC<{ children: ReactNode }> = ({
  children,
}: AppLayoutProps) => {
  return (
    <div className="flex pt-[130px] w-full justify-center px-4 md:px-6">
      <div className="w-full lg:w-[856px] md:w-[641px]">{children}</div>
    </div>
  );
};

export default AppLayout;
