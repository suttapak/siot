import React from 'react';
import io from 'socket.io-client';

export interface SocketIOTpye {
  client: SocketIOClient.Socket;
}

const socketIOContext = React.createContext<SocketIOTpye>({} as SocketIOTpye);

interface SocketIOProviderProps {
  children: React.ReactNode;
}
export const ProviderSokcetIO: React.FC<SocketIOProviderProps> = ({ children }) => {
  const toast = useSocketIoProvider();
  return <socketIOContext.Provider value={toast}>{children}</socketIOContext.Provider>;
};

export const useSocketIO = () => {
  return React.useContext(socketIOContext);
};

var socket = io('ws://localhost:4000/', {
  reconnectionDelayMax: 10000,
  transports: ['websocket'],
});

const useSocketIoProvider: () => SocketIOTpye = () => {
  return { client: socket };
};
