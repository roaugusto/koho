import React, { useState, useEffect } from 'react';
import moment from 'moment';

import { FiDownload } from 'react-icons/fi';

// import Header from '../../components/Header';

import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';
import Tooltip from '@material-ui/core/Tooltip';
import { Container, ButtonDownload } from './styles';
import Navbar from '../../components/Navbar';
import api from '../../services/api';
import { useToast } from '../../hooks/toast';

interface Record {
  id: string;
  customer_id: string;
  load_amount: string;
  time: Date;
  accepted: string;
  codError: string;
  message: string;
  formatted_date: string;
}

const Dashboard: React.FC = () => {
  const { addToast } = useToast();
  const [records, setRecords] = useState<Record[]>([]);
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);

  const useStyles = makeStyles({
    root: {
      width: '100%',
    },
    container: {
      maxHeight: 300,
    },
  });

  const classes = useStyles();

  useEffect(() => {
    async function loadRecords(): Promise<void> {
      const lastUUIDFile = localStorage.getItem('@Koho:LastUUIDFile');

      if (lastUUIDFile) {
        api
          .get(`/api/funds/result?process_id=${lastUUIDFile}`)
          .then(response => {
            if (response.data) {
              const recordsResponse = response.data.map((record: Record) => ({
                ...record,
                formatted_date: moment
                  .utc(record.time)
                  .format('DD/MM/yyyy HH:mm:ss'),
              }));
              setRecords(recordsResponse);
            }
          });
      }
    }

    loadRecords();
  }, []);

  const handleChangePage = (event: any, newPage: number): void => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event: any): void => {
    setRowsPerPage(+event.target.value);
    setPage(0);
  };

  const handleDownload = async (): Promise<void> => {
    const lastUUIDFile = localStorage.getItem('@Koho:LastUUIDFile');
    const config = {
      headers: { 'Cache-Control': 'no-cache' },
    };

    try {
      const res = await api.get(
        `/api/funds/download?uuid_file=${lastUUIDFile}`,
        config,
      );
      const url = window.URL.createObjectURL(new Blob([res.data]));
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', 'output.txt'); // or any other extension
      document.body.appendChild(link);
      link.click();
      link.remove();
    } catch (err) {
      addToast({
        type: 'error',
        title: 'Download error',
        description: 'Output file not found. Please import a new file again!',
      });
    }
  };

  return (
    <>
      <Navbar />
      <Container>
        <Paper className={classes.root}>
          <TableContainer>
            <Table stickyHeader aria-label="sticky table">
              <TableHead>
                <TableRow>
                  <TableCell align="right" style={{ width: 70 }}>
                    Transaction ID
                  </TableCell>
                  <TableCell align="right" style={{ width: 70 }}>
                    Customer ID
                  </TableCell>
                  <TableCell align="right" style={{ width: 80 }}>
                    Load Amount
                  </TableCell>
                  <TableCell align="center" style={{ width: 100 }}>
                    Date / Time
                  </TableCell>
                  <TableCell align="center" style={{ width: 50 }}>
                    Accepted
                  </TableCell>
                  <TableCell style={{ width: 250 }}>Message</TableCell>
                  <TableCell style={{ width: 10 }}>
                    <Tooltip title="Dowload last result" placement="top">
                      <ButtonDownload onClick={handleDownload} type="button">
                        <FiDownload />
                      </ButtonDownload>
                    </Tooltip>
                  </TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {records
                  .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                  .map((record, key) => {
                    return (
                      <TableRow hover role="checkbox" tabIndex={-1} key={key}>
                        <TableCell align="right">{record.id}</TableCell>
                        <TableCell align="right">
                          {record.customer_id}
                        </TableCell>
                        <TableCell align="right">
                          {record.load_amount}
                        </TableCell>
                        <TableCell align="center">
                          {record.formatted_date}
                        </TableCell>
                        <TableCell align="center">{record.accepted}</TableCell>
                        <TableCell>{record.message}</TableCell>
                        <TableCell style={{ width: 10 }} />
                      </TableRow>
                    );
                  })}
              </TableBody>
            </Table>
          </TableContainer>
          <TablePagination
            rowsPerPageOptions={[10, 25, 100]}
            component="div"
            count={records.length}
            rowsPerPage={rowsPerPage}
            page={page}
            onChangePage={handleChangePage}
            onChangeRowsPerPage={handleChangeRowsPerPage}
          />
        </Paper>
      </Container>
    </>
  );
};

export default Dashboard;
