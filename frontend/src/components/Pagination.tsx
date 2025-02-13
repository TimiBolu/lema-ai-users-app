import DirectionBtn from "./DirectionBtn";

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

export default Pagination;
