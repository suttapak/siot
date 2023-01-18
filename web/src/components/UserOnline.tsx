import React, { useRef } from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';
import { Box, Paper, Typography } from '@mui/material';
import { useSocketIO } from '../hooks/useSocketIO';
import { UserOnlineType } from '../types/UserOnlineType';
import moment from 'moment';

type Props = {};

const UserOnline = (props: Props) => {
  const { client } = useSocketIO();

  const [data, setData] = React.useState<[] | UserOnlineType[]>([]);

  React.useEffect(() => {
    client.emit('userOnline');
  }, []);

  const refWidth = useRef<HTMLDivElement>(null);

  const [width, setWidth] = React.useState(300);

  React.useEffect(() => {
    setWidth(refWidth.current?.clientWidth ? refWidth.current?.clientWidth : 300);
  }, [refWidth, refWidth.current?.clientWidth]);

  client.on('userOnline', (value: UserOnlineType[]) => {
    setData(
      value
        .map((v) => {
          const label = moment(v.createdAt).format('h:mm Do MMMM');
          v = {
            id: v.id,
            createdAt: v.createdAt,
            updatedAt: v.updatedAt,
            onLineCount: v.onLineCount,
            label: label,
          };
          return v;
        })
        .sort((a, b) => a.id - b.id)
    );
  });
  return (
    <React.Fragment>
      <Paper
        elevation={0}
        sx={{
          height: 360,
          marginTop: 2,
        }}
      >
        <Box
          sx={{
            padding: 1,
          }}
        >
          <Typography variant='h6'>ผู้ใช้งาน</Typography>
        </Box>
        <Box
          ref={refWidth}
          sx={{
            padding: 1,
          }}
        >
          <LineChart
            width={width - 10}
            height={264.2}
            data={data}
            margin={{
              top: 5,
              right: 30,
              left: 20,
              bottom: 5,
            }}
          >
            <CartesianGrid strokeDasharray='3 3' />
            <XAxis dataKey={'label'} />
            <YAxis />
            <Tooltip />
            <Line type='monotone' dataKey='onLineCount' stroke='#8884d8' dot={<div></div>} />
          </LineChart>
        </Box>
      </Paper>
    </React.Fragment>
  );
};

export default UserOnline;
