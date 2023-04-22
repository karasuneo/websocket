import { useRecoilCallback, useRecoilValue } from "recoil";

import { websocketAtom } from "@/store/websocket";
import { messageListAtom } from "@/store/message";

import { Message } from "@/types/message";
import { IMessageEvent } from "websocket";

export const useMessageList = (): Message[] => {
  const socket = useRecoilValue(websocketAtom);
  const messageList = useRecoilValue(messageListAtom);
  var tempProps = JSON.parse(JSON.stringify(socket));

  const updateMessageList = useRecoilCallback(
    ({ set }) =>
      (message: Message) => {
        set(messageListAtom, [...messageList, message]);
      }
  );
  tempProps.onmessage = (msg: IMessageEvent) => {
    var data = msg.data
    const content = JSON.parse(data as string);
    const message: Message = { content: content };
    updateMessageList(message);
  };

  return messageList;
};
