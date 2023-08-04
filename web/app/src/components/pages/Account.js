import React, { useState, useEffect } from 'react';
import { get, postFormData } from './../../services/apiClient';

const AccountPage = ({ user }) => {
    const [uploadedFiles, setUploadedFiles] = useState([]);
    const [fileList, setFileList] = useState([]);
    const [page, setPage] = useState(0);
    const limit = 10; // Set the number of files per page

    // Function to handle file drop event
    const handleFileDrop = (event) => {
        event.preventDefault();
        const incomingFile = event.dataTransfer.files[0];
        if (incomingFile) {
            setUploadedFiles([incomingFile]);
        }
    };

    // Function to handle file upload
    const handleFileUpload = async () => {
        try {
            if (uploadedFiles.length === 0) {
                console.error('No file selected for upload.');
                return;
            }

            const formData = new FormData();
            formData.append('file', uploadedFiles[0]);

            // Make a post request using postFormData function
            const result = await postFormData('/file', formData);
            console.log('File upload response:', result);

            // Clear dropped files after successful upload
            setUploadedFiles([]);
            // Fetch the updated list of files after upload
            fetchFiles();
        } catch (error) {
            console.error('Error uploading files:', error);
        }
    };

    // Function to fetch the list of files from the backend
    const fetchFiles = async () => {
        try {
            const response = await get(`/file/list/all?limit=${limit}&page=${page}`);
            if (response.success) {
                if (response.files.length === null) {
                    setFileList([]);
                } else {
                    setFileList(response.files);
                }
            }
        } catch (error) {
            console.error('Error fetching files:', error);
        }
    };

    // Call the fetchFiles function when the component loads or when the page changes
    useEffect(() => {
        fetchFiles();
    }, [page]);

    return (
        <div>
            <h2>Account Details</h2>
            <p>Username: {user.username}</p>
            <p>Email: {user.email}</p>

            {/* Drop zone */}
            <div
                style={{
                    border: '2px dashed #ccc',
                    padding: '20px',
                    margin: '20px 0',
                    textAlign: 'center',
                }}
                onDrop={handleFileDrop}
                onDragOver={(e) => e.preventDefault()}
            >
                <p>Drag and drop files here</p>
                <button onClick={handleFileUpload}>Upload Files</button>
                {uploadedFiles.map((file, index) => (
                    <p key={index}>{file.name}</p>
                ))}
            </div>

            {/* List of uploaded files */}
            <div>
                {fileList.map((file) => (
                    <p key={file.id}>{file.filename}</p>
                ))}
            </div>

            {/* Pagination */}
            <div>
                <button
                    onClick={() => setPage((prevPage) => Math.max(prevPage - 1, 0))}
                    disabled={page === 0}
                >
                    Previous Page
                </button>
                <button onClick={() => setPage((prevPage) => prevPage + 1)}>Next Page</button>
            </div>
        </div>
    );
};

export default AccountPage;
