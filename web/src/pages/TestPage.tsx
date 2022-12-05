// @ts-nocheck

import React, { useState, useEffect } from 'react';
import io from 'socket.io-client';

var socket = io('ws://localhost:4000/', {
  reconnectionDelayMax: 10000,
  transports: ['websocket'],
});

function TestPage() {
  const [isConnected, setIsConnected] = useState(socket.connected);
  const [lastPong, setLastPong] = useState<string | null>(null);

  useEffect(() => {
    socket.on('connect', () => {
      socket.emit('subscript', { boxId: '0384d656-0c9b-4c0c-8c6d-26004d855aca' });
    });

    socket.on('betamanga-0384d656', (data) => {
      console.log(data);

      setIsConnected(false);
    });

    return () => {
      socket.off('connect');
      socket.off('disconnect');
      socket.off('pong');
    };
  }, []);

  return (
    <div>
      <p>Connected: {'' + isConnected}</p>
      <p>Last pong: {lastPong || '-'}</p>
      <button>Send ping</button>
    </div>
  );
}

export default TestPage;
