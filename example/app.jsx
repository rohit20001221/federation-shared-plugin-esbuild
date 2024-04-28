import { useState } from "react";

let Greet = () => {
  const [count, setCount] = useState(0);

  return (
    <div>
      <div>{count}</div>
      <div onClick={() => setCount((c) => c + 1)}>
        <button>click</button>
      </div>
      <h1>Hello, world!</h1>
    </div>
  );
};

export default Greet;
