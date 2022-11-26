import React from 'react';

export interface toastContextInterface {
  setMessage: React.Dispatch<React.SetStateAction<string[]>>;
  message: string[];
  addMessage: (msg: string) => void;
  removeMessage: (index: number) => void;
}

const toastContext = React.createContext<toastContextInterface>(
  {} as toastContextInterface
);

interface Props {
  children: React.ReactNode;
}
export const ProviderToast: React.FC<Props> = ({ children }) => {
  const toast = useProviderToast();
  return (
    <toastContext.Provider value={toast}>{children}</toastContext.Provider>
  );
};

export const useToast = () => {
  return React.useContext(toastContext);
};

const useProviderToast: () => toastContextInterface = () => {
  const [message, setMessage] = React.useState<string[]>([]);
  function addMessage(msg: string) {
    setMessage((prev) => [...prev, msg]);
  }

  function removeMessage(index: number) {
    setMessage((prev) => prev.filter((_, i) => index !== i));
  }
  return { message, setMessage, addMessage, removeMessage };
};
