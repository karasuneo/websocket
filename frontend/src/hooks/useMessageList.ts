import { useRecoilCallback, useRecoilValue } from "recoil";

import { websocketAtom } from "@/store/websocket";
import { messageListAtom } from "@/store/message";

import { Message } from "@/types/message";

export const useMessageList = (): Message[] => {
  const socket = useRecoilValue(websocketAtom);
  const messageList = useRecoilValue(messageListAtom);

  const updateMessageList = useRecoilCallback(
    ({ set }) =>
      (message: Message) => {
        set(messageListAtom, [...messageList, message]);
      }
  );
  socket.onmessage = (msg) => {
    const content = JSON.parse(msg.data as string);
    const message: Message = { content: content };
    updateMessageList(message);
  };

  return messageList;
};
