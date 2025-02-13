import { useEffect, useState } from "react";
import { useNavigate } from "@tanstack/react-router";

import Loader from "../components/Loader";

import { useWindowSize } from "../hooks/useWindowSize";
import ErrorMessage from "../components/ErrorMgr/ErrorMessage";
import Pagination from "../components/Pagination";
import { useUsers } from "../hooks/useUsers";

const UsersTable = () => {
  const [page, setPage] = useState(1);
  const { isSmallerScreen } = useWindowSize();
  const [totalPages, setTotalPages] = useState<number | null>(null);

  const navigate = useNavigate({ from: "/" });
  const { isPending, error, data } = useUsers(page);

  useEffect(() => {
    if (data?.data?.pagination?.totalPages) {
      setTotalPages(data?.data.pagination.totalPages);
    }
  }, [data?.data?.pagination?.totalPages]);

  const pages = totalPages ?? 1;
  if (error) return <ErrorMessage error={error!} />;

  return (
    <>
      <p className="text-[#181D27] text-[60px]/[72px]">Users</p>
      <div className="min-w-full mt-[24px] h-[323px] border border-[#E9EAEB] rounded-lg overflow-hidden text-[#535862]">
        {isPending ? (
          <div className="flex w-full h-full justify-center items-center">
            <Loader color="#bdb1c6" size={50} />
          </div>
        ) : (
          <div className="w-full overflow-x-auto">
            <table className="w-full">
              <thead>
                <tr>
                  <th className="py-3 pl-6 text-left text-xs font-medium">
                    Full Name
                  </th>
                  <th className="py-3 pl-6 text-left text-xs font-medium">
                    Email
                  </th>
                  <th className="py-3 pl-6 text-left text-xs font-medium w-[392px]">
                    Address
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                {data?.data?.users?.map((user) => (
                  <tr
                    key={user.id}
                    className="hover:bg-gray-50 cursor-pointer"
                    onClick={() => navigate({ to: `/users/${user.id}/posts` })}
                  >
                    <td
                      className="px-6 py-6 text-sm font-medium truncate overflow-hidden whitespace-nowrap
                          w-[124px] md:w-[124px] lg:w-[200px]"
                    >
                      {`${user.firstname} ${user.lastname}`}
                    </td>

                    <td
                      className="px-6 py-6 text-sm font-normal truncate overflow-hidden whitespace-nowrap
                          w-[124px] md:w-[124px] lg:w-[264px]"
                    >
                      {user.email}
                    </td>

                    <td className="px-6 py-6 text-sm font-normal truncate w-[392px]">
                      {`${user.address.street}, ${user.address.city}, ${user.address.state} ${user.address.zipCode}`}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>

      {/* Pagination */}
      <div className="flex w-full justify-end mt-4">
        <Pagination
          page={page}
          setPage={setPage}
          totalPages={pages}
          isSmallerScreen={isSmallerScreen}
        />
      </div>
    </>
  );
};

export default UsersTable;
