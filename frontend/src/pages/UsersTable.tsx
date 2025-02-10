import { useEffect, useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";

import Loader from "../components/Loader";
import DirectionBtn from "../components/DirectionBtn";

import { fetchUsers } from "../apis";
import { useWindowSize } from "../hooks/useWindowSize";
import ErrorMessage from "../components/ErrorMgr/ErrorMessage";

const Pagination = ({
  page,
  setPage,
  totalPages,
  isSmallerScreen,
}: {
  page: number;
  totalPages: number;
  isSmallerScreen: boolean;
  setPage: React.Dispatch<React.SetStateAction<number>>;
}) => {
  const getPageNumbers = () => {
    const maxPages = isSmallerScreen ? 4 : 6;
    const maxPagePortion = isSmallerScreen ? 2 : 3;
    const pageList = Array.from(Array(totalPages), (_, i) => i + 1);
    if (totalPages <= maxPages) {
      return pageList.map((_, i) => i + 1);
    }

    // where X = totalPages;
    // is the page within the range<...> or on either ends?
    // If it is within the range then render this view 1 2 ... <page> ... X-1 X
    // else render the extended view 1 2 3 ... X-2 X-1 X
    const withinRange =
      page > maxPagePortion && page <= totalPages - maxPagePortion;
    return withinRange
      ? [
          1,
          ...(!isSmallerScreen ? [2] : []),
          "...",
          page,
          "...",
          ...(!isSmallerScreen ? [totalPages - 1] : []),
          totalPages,
        ]
      : [
          ...pageList.slice(0, maxPagePortion),
          "...",
          ...pageList.slice(totalPages - maxPagePortion, totalPages),
        ];
  };

  const numberList = getPageNumbers();

  return (
    <div className="flex gap-[42px] items-center">
      <DirectionBtn
        direction="back"
        disabled={page === 1}
        onClick={() => setPage((prev) => Math.max(prev - 1, 1))}
      />

      <div className="w-full flex space-x-2">
        {numberList.map((num, index) =>
          num === "..." ? (
            <span key={`ellipsis-${index}`} className="px-3 py-1 text-gray-500">
              ...
            </span>
          ) : (
            <button
              key={`page-${num}`}
              aria-label={`Page ${num}`}
              onClick={() => setPage(num as number)}
              className={`cursor-pointer px-3 py-1 ${
                page === num
                  ? "bg-[#F9F5FF] text-[#7F56D9] rounded"
                  : "bg-white text-[#717680]"
              }`}
            >
              {num}
            </button>
          ),
        )}
      </div>

      <DirectionBtn
        direction="next"
        disabled={page === totalPages}
        onClick={() => setPage((prev) => Math.min(prev + 1, totalPages))}
      />
    </div>
  );
};

const UsersTable = () => {
  const [page, setPage] = useState(1);
  const { isSmallerScreen } = useWindowSize();
  const [totalPages, setTotalPages] = useState<number | null>(null);

  const navigate = useNavigate({ from: "/" });
  const { isPending, error, data } = useQuery({
    queryKey: ["users", page],
    queryFn: () => fetchUsers(page),
  });

  useEffect(() => {
    if (data?.pagination?.totalPages) {
      setTotalPages(data.pagination.totalPages);
    }
  }, [data?.pagination?.totalPages]);

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
                {data?.users?.map((user) => (
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
