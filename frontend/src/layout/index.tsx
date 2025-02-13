import { FC, ReactNode } from "react";

type AppLayoutProps = {
  children: ReactNode;
};

const AppLayout: FC<AppLayoutProps> = ({ children }) => {
  return (
    <div className="flex pt-[130px] w-full justify-center px-4 md:px-6">
      <div className="w-full max-w-[calc(100vw-32px)] lg:max-w-[856px]">
        {children}
      </div>
    </div>
  );
};

export default AppLayout;
