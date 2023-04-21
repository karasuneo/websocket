import { Inter } from "next/font/google";
import { NextPage } from "next";
import { MessageInput } from "@/components/MessageInput";
import { MessageList } from "@/components/MessageList";

const inter = Inter({ subsets: ["latin"] });

const Home: NextPage = () => {
  return (
    <>
      <h1>Simple Chat</h1>
      <MessageInput />
      <MessageList />
    </>
  );
};

export default Home;

// import { Inter } from "next/font/google";
// import { useEffect, useRef, useState } from "react";
// import { NextPage } from "next";

// const inter = Inter({ subsets: ["latin"] });

// const Home: NextPage = () => {
//   const socketRef = useRef<WebSocket>();
//   const [isConnected, setIsConnected] = useState(false);

//   useEffect(() => {
//     socketRef.current = new WebSocket("ws://localhost:8080/socket");
//     socketRef.current.onopen = function () {
//       setIsConnected(true);
//       console.log("Connected");
//     };

//     socketRef.current.onclose = function () {
//       console.log("closed");
//       setIsConnected(false);
//     };

//     return () => {
//       if (socketRef.current == null) {
//         return;
//       }
//       socketRef.current.close();
//     };
//   }, []);

//   return (
//     <>
//       <h1>WebSocket is connected : {`${isConnected}`}</h1>
//     </>
//   );
// };

// export default Home;
