import { ListItem, Typography } from "@mui/material";
import ReactMarkdown from "react-markdown";
import rehypeHighlight from "rehype-highlight";
import rehypeRaw from "rehype-raw";
import remarkGfm from "remark-gfm";

export interface Message {
  role: string;
  content: string;
}

type MessageProps = {
  message: Message;
};

export function MessageBlock(props: MessageProps) {
  return (
    <ListItem
      disablePadding
      sx={{
        backgroundColor: (theme) =>
          props.message.role == "assistant"
            ? theme.palette.grey[300]
            : theme.palette.grey[100],
      }}
    >
      <Typography
        sx={{
          width: "100%",
          maxWidth: "100%",
          paddingX: "12%",
          paddingY: "3%",
        }}
      >
        {props.message.role == "assistant" ? (
          <ReactMarkdown
            className={"markdown-body"}
            children={props.message.content}
            remarkPlugins={[remarkGfm]}
            rehypePlugins={[rehypeRaw, rehypeHighlight]}
          />
        ) : (
          <div>{props.message.content}</div>
        )}
      </Typography>
    </ListItem>
  );
}
