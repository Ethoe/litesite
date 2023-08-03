import React, { useState } from 'react';
import { postFormData } from './../../services/apiClient';


const AccountPage = ({ user }) => {
    const [files, setFiles] = useState([]);

    // Function to handle file drop event
    const handleFileDrop = (event) => {
        event.preventDefault();
        const incomingFile = event.dataTransfer.files[0];
        if (incomingFile) {
            setFiles([incomingFile]);
        }
    };

    // Function to handle file upload
    const handleFileUpload = async () => {
        try {
            if (files.length === 0) {
                console.error('No file selected for upload.');
                return;
            }

            const formData = new FormData();
            formData.append('file', files[0]);

            // Make a post request using postFormData function
            const result = await postFormData('/file', formData);
            console.log('File upload response:', result);

            // Clear dropped files after successful upload
            setFiles([]);
        } catch (error) {
            console.error('Error uploading files:', error);
        }
    };

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
                {files.map((file, index) => (
                    <p key={index}>{file.name}</p>
                ))}
            </div>
        </div>
    );
};

export default AccountPage;
