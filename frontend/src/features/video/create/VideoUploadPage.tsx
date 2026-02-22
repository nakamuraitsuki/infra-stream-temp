import { useNavigate } from "react-router"
import { useCreateVideo } from "@/features/video/create/use-create-video";
import { useState } from "react";
import type { VideoTag } from "../../../domain/video/video.model";

export const VideoUploadPage = () => {
  const navigate = useNavigate();
  const { createVideo, createLoading, uploadLoading } = useCreateVideo();

  const [title, setTitle] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [tags, _setTags] = useState<VideoTag[]>([]);
  const [file, setFile] = useState<File | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files?.[0];
    if (selectedFile) {
      setFile(selectedFile);
    }
  };

  const handleSubmit = async (e: React.SyntheticEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!file || !title) return;

    setError(null);

    const res = await createVideo(
      title,
      description,
      tags,
      file,
    );

    if (res.type === "success") {
      navigate("/my-page");
    } else {
      setError(`${res.type}: ${res.error}`);
    }
  };

  return (
    <div style={{ padding: "20px", maxWidth: "600px", margin: "0 auto" }}>
      <button onClick={() => navigate(-1)}>Back</button>
      <h2>Upload New Video</h2>

      <form onSubmit={handleSubmit} style={{ display: "flex", flexDirection: "column", gap: "15px" }}>
        <div>
          <label htmlFor="title" style={{ display: "block" }}>Title</label>
          <input
            id="title"
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
            style={{ width: "100%" }}
          />
        </div>

        <div>
          <label htmlFor="description" style={{ display: "block" }}>Description</label>
          <textarea
            id="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            style={{ width: "100%", height: "100px" }}
          />
        </div>

        <div>
          <label htmlFor="videoFile" style={{ display: "block" }}>Video File</label>
          <input
            id="videoFile"
            type="file"
            accept="video/*"
            onChange={handleFileChange}
            required
          />
        </div>

        {error && <p style={{ color: "red" }}>{error}</p>}

        <button type="submit" disabled={createLoading || uploadLoading}>
          {createLoading ? "Creating Metadata..." :
            uploadLoading ? "Uploading Video File..." :
              "Start Upload"}
        </button>
      </form>

      {(createLoading || uploadLoading) && (
        <div style={{ marginTop: "10px", fontSize: "0.9em", color: "#666" }}>
          ※ 処理を中断しないでください
        </div>
      )}
    </div>
  )
}