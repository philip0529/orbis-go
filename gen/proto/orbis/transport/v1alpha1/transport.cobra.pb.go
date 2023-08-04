// Code generated by protoc-gen-cobra. DO NOT EDIT.

package transportv1alpha1

import (
	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	flag "github.com/NathanBaulch/protoc-gen-cobra/flag"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	pb "github.com/libp2p/go-libp2p/core/crypto/pb"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func TransportServiceClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("TransportService"),
		Short: "TransportService service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_TransportServiceGetHostCommand(cfg),
		_TransportServiceConnectCommand(cfg),
		_TransportServiceSendCommand(cfg),
		_TransportServiceGossipCommand(cfg),
		_TransportServiceNewMessageCommand(cfg),
	)
	return cmd
}

func _TransportServiceGetHostCommand(cfg *client.Config) *cobra.Command {
	req := &GetHostRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetHost"),
		Short: "GetHost RPC client",
		Long:  "GetHost returns the information about the host node.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "TransportService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "TransportService", "GetHost"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransportServiceClient(cc)
				v := &GetHostRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetHost(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Transport, cfg.FlagNamer("Transport"), "", "")

	return cmd
}

func _TransportServiceConnectCommand(cfg *client.Config) *cobra.Command {
	req := &ConnectRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Connect"),
		Short: "Connect RPC client",
		Long:  "Connect connects to a peer node.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "TransportService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "TransportService", "Connect"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransportServiceClient(cc)
				v := &ConnectRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Connect(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Transport, cfg.FlagNamer("Transport"), "", "")
	cmd.PersistentFlags().StringVar(&req.Id, cfg.FlagNamer("Id"), "", "")
	cmd.PersistentFlags().StringVar(&req.Address, cfg.FlagNamer("Address"), "", "")

	return cmd
}

func _TransportServiceSendCommand(cfg *client.Config) *cobra.Command {
	req := &SendRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Send"),
		Short: "Send RPC client",
		Long:  "Send sends a message to a peer node.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "TransportService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "TransportService", "Send"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransportServiceClient(cc)
				v := &SendRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Send(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	_Node := &Node{}
	cmd.PersistentFlags().StringVar(&_Node.Id, cfg.FlagNamer("Node Id"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Node Id"), func() { req.Node = _Node })
	cmd.PersistentFlags().StringVar(&_Node.Address, cfg.FlagNamer("Node Address"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Node Address"), func() { req.Node = _Node })
	_Node_PublicKey := &pb.PublicKey{}
	flag.EnumPointerVar(cmd.PersistentFlags(), &_Node_PublicKey.Type, cfg.FlagNamer("Node PublicKey Type"), "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Node PublicKey Type"), func() { req.Node = _Node; _Node.PublicKey = _Node_PublicKey })
	flag.BytesBase64Var(cmd.PersistentFlags(), &_Node_PublicKey.Data, cfg.FlagNamer("Node PublicKey Data"), "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Node PublicKey Data"), func() { req.Node = _Node; _Node.PublicKey = _Node_PublicKey })
	_Message := &Message{}
	cmd.PersistentFlags().Int64Var(&_Message.Timestamp, cfg.FlagNamer("Message Timestamp"), 0, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Timestamp"), func() { req.Message = _Message })
	cmd.PersistentFlags().StringVar(&_Message.Id, cfg.FlagNamer("Message Id"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Id"), func() { req.Message = _Message })
	cmd.PersistentFlags().StringVar(&_Message.Type, cfg.FlagNamer("Message Type"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Type"), func() { req.Message = _Message })
	flag.BytesBase64Var(cmd.PersistentFlags(), &_Message.Payload, cfg.FlagNamer("Message Payload"), "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Payload"), func() { req.Message = _Message })
	cmd.PersistentFlags().BoolVar(&_Message.Gossip, cfg.FlagNamer("Message Gossip"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Gossip"), func() { req.Message = _Message })
	cmd.PersistentFlags().StringVar(&_Message.NodeId, cfg.FlagNamer("Message NodeId"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message NodeId"), func() { req.Message = _Message })
	flag.BytesBase64Var(cmd.PersistentFlags(), &_Message.NodePubKey, cfg.FlagNamer("Message NodePubKey"), "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message NodePubKey"), func() { req.Message = _Message })
	cmd.PersistentFlags().StringVar(&_Message.RingId, cfg.FlagNamer("Message RingId"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message RingId"), func() { req.Message = _Message })
	flag.BytesBase64Var(cmd.PersistentFlags(), &_Message.Signature, cfg.FlagNamer("Message Signature"), "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Signature"), func() { req.Message = _Message })
	cmd.PersistentFlags().StringVar(&_Message.TargetId, cfg.FlagNamer("Message TargetId"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message TargetId"), func() { req.Message = _Message })
	flag.BytesBase64Var(cmd.PersistentFlags(), &_Message.TargetPubKey, cfg.FlagNamer("Message TargetPubKey"), "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message TargetPubKey"), func() { req.Message = _Message })

	return cmd
}

func _TransportServiceGossipCommand(cfg *client.Config) *cobra.Command {
	req := &GossipRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Gossip"),
		Short: "Gossip RPC client",
		Long:  "Gossip broadcasts a message to a topic.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "TransportService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "TransportService", "Gossip"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransportServiceClient(cc)
				v := &GossipRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Gossip(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Transport, cfg.FlagNamer("Transport"), "", "")
	cmd.PersistentFlags().StringVar(&req.Topic, cfg.FlagNamer("Topic"), "", "")
	_Message := &Message{}
	cmd.PersistentFlags().Int64Var(&_Message.Timestamp, cfg.FlagNamer("Message Timestamp"), 0, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Timestamp"), func() { req.Message = _Message })
	cmd.PersistentFlags().StringVar(&_Message.Id, cfg.FlagNamer("Message Id"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Id"), func() { req.Message = _Message })
	cmd.PersistentFlags().StringVar(&_Message.Type, cfg.FlagNamer("Message Type"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Type"), func() { req.Message = _Message })
	flag.BytesBase64Var(cmd.PersistentFlags(), &_Message.Payload, cfg.FlagNamer("Message Payload"), "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Payload"), func() { req.Message = _Message })
	cmd.PersistentFlags().BoolVar(&_Message.Gossip, cfg.FlagNamer("Message Gossip"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Gossip"), func() { req.Message = _Message })
	cmd.PersistentFlags().StringVar(&_Message.NodeId, cfg.FlagNamer("Message NodeId"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message NodeId"), func() { req.Message = _Message })
	flag.BytesBase64Var(cmd.PersistentFlags(), &_Message.NodePubKey, cfg.FlagNamer("Message NodePubKey"), "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message NodePubKey"), func() { req.Message = _Message })
	cmd.PersistentFlags().StringVar(&_Message.RingId, cfg.FlagNamer("Message RingId"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message RingId"), func() { req.Message = _Message })
	flag.BytesBase64Var(cmd.PersistentFlags(), &_Message.Signature, cfg.FlagNamer("Message Signature"), "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message Signature"), func() { req.Message = _Message })
	cmd.PersistentFlags().StringVar(&_Message.TargetId, cfg.FlagNamer("Message TargetId"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message TargetId"), func() { req.Message = _Message })
	flag.BytesBase64Var(cmd.PersistentFlags(), &_Message.TargetPubKey, cfg.FlagNamer("Message TargetPubKey"), "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Message TargetPubKey"), func() { req.Message = _Message })

	return cmd
}

func _TransportServiceNewMessageCommand(cfg *client.Config) *cobra.Command {
	req := &NewMessageRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("NewMessage"),
		Short: "NewMessage RPC client",
		Long:  "NewMessage returns a formated messages.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "TransportService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "TransportService", "NewMessage"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransportServiceClient(cc)
				v := &NewMessageRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.NewMessage(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Transport, cfg.FlagNamer("Transport"), "", "")
	cmd.PersistentFlags().StringVar(&req.Id, cfg.FlagNamer("Id"), "", "")
	cmd.PersistentFlags().BoolVar(&req.Gossip, cfg.FlagNamer("Gossip"), false, "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Payload, cfg.FlagNamer("Payload"), "")
	cmd.PersistentFlags().StringVar(&req.Type, cfg.FlagNamer("Type"), "", "")
	cmd.PersistentFlags().StringVar(&req.RingId, cfg.FlagNamer("RingId"), "", "")

	return cmd
}
