import { Message } from "@/types/message";
import { atom } from "recoil";

export const messageListAtom = atom<Message[]>({
  key: "messageList",
  default: [],
});
