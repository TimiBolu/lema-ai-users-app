type MockAPIResponseProps<T> = {
  success?: boolean;
  message?: string;
  data: T;
};
const MockAPIResponse = <T>({
  success = true,
  message = "successful",
  data,
}: MockAPIResponseProps<T>) => {
  return {
    success,
    message,
    data,
  };
};

export default MockAPIResponse;
