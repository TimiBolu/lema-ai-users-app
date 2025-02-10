const ErrorMessage = ({ error }: { error: Error }) => (
  <div className="rounded-lg bg-red-50 p-4">
    <h3 className="text-red-800">Ooops! Something went wrong</h3>
    <p className="text-red-600">{error?.message}</p>
    <button className="underline" onClick={() => window.location.reload()}>
      Try Again
    </button>
  </div>
);

export default ErrorMessage;
