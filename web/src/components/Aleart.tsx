interface AleartProps {
  type: 'error' | 'worning';
  message?: string;
}

export function Alert(props: AleartProps) {
  return (
    <>
      <div
        className={`${
          props.type === 'error' ? 'bg-red-500 text-white' : 'bg-yellow-500 '
        } w-full  rounded  px-4 py-1 border border-red-600 shadow drop-shadow flex justify-center items-center text-md md:text-lg font-bold`}
      >
        {props.message ? props.message : 'unknow ' + props.type + '.'}
      </div>
    </>
  );
}
