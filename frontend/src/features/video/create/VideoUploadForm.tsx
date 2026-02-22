import { useState } from "react";
import type { VideoTag } from "@/domain/video/video.model";
import { useCreateVideo } from "./use-create-video";

type Props = {
  onSuccess: () => void;
  onBack: () => void;
};

export const VideoUploadForm = ({ onSuccess, onBack }: Props) => {
  const { createVideo, createLoading, uploadLoading } = useCreateVideo();

  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [tags] = useState<VideoTag[]>([]);
  const [file, setFile] = useState<File | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.SyntheticEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!file || !title) return;

    setError(null);

    const res = await createVideo(
      title,
      description,
      tags,
      file
    );

    if (res.type === "success") {
      onSuccess();
    } else {
      setError(`${res.type}: ${res.error}`);
    }
  };

  return (
    <div style={{ padding: "20px", maxWidth: "600px", margin: "0 auto" }}>
      <button onClick={onBack}>Back</button>

      <h2>Upload New Video</h2>

      <form
        onSubmit={handleSubmit}
        style={{ display: "flex", flexDirection: "column", gap: "15px" }}
      >
        <div>
          <label htmlFor="title">Title</label>
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
          <label htmlFor="description">Description</label>
          <textarea
            id="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            style={{ width: "100%", height: "100px" }}
          />
        </div>

        <div>
          <label htmlFor="videoFile">Video File</label>
          <input
            id="videoFile"
            type="file"
            accept="video/*"
            required
            onChange={(e) => {
              const selected = e.target.files?.[0];
              if (selected) setFile(selected);
            }}
          />
        </div>

        {error && <p style={{ color: "red" }}>{error}</p>}

        <button type="submit" disabled={createLoading || uploadLoading}>
          {createLoading
            ? "Creating Metadata..."
            : uploadLoading
            ? "Uploading Video File..."
            : "Start Upload"}
        </button>
      </form>

      {(createLoading || uploadLoading) && (
        <div style={{ marginTop: "10px", fontSize: "0.9em", color: "#666" }}>
          ※ 処理を中断しないでください
        </div>
      )}
    </div>
  );
};
