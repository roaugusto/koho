import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';

import filesize from 'filesize';

import Navbar from '../../components/Navbar';
import Loader from '../../components/Loader';

import FileList from '../../components/FileList';
import Upload from '../../components/Upload';

import { Container, Title, ImportFileContainer, Footer } from './styles';

import api from '../../services/api';

interface FileProps {
  file: File;
  name: string;
  readableSize: string;
}

export interface IUpload {
  process_id: string;
}

const Import: React.FC = () => {
  const [showPopup, setShowPopup] = useState(false);
  const [uploadedFiles, setUploadedFiles] = useState<FileProps[]>([]);
  const history = useHistory();

  async function handleUpload(): Promise<void> {
    if (!uploadedFiles.length) return;

    const promises = uploadedFiles.map(async item => {
      try {
        setShowPopup(true);
        const data = new FormData();
        data.append('file', item.file, item.name);
        const res = await api.post<IUpload>('/api/funds-write-result-db', data);
        const uuid = res.data.process_id
        localStorage.setItem('@Koho:LastUUIDFile', uuid);
        setShowPopup(false);
      } catch (err) {
        console.log(err);
      }
    });

    await Promise.all(promises);

    history.goBack();
  }

  function submitFile(files: File[]): void {
    const listFiles = files.map(file => ({
      file,
      name: file.name,
      readableSize: filesize(file.size),
    }));

    setUploadedFiles(listFiles);
  }

  return (
    <>
      <Navbar />
      <Container>
        <Title>Import a file</Title>
        <ImportFileContainer>
          <Upload onUpload={submitFile} />
          {!!uploadedFiles.length && <FileList files={uploadedFiles} />}

          <Footer>
            <button onClick={handleUpload} type="button">
              Send
            </button>
          </Footer>
        </ImportFileContainer>
      </Container>
      <Loader showPopup={showPopup} />
    </>
  );
};

export default Import;
