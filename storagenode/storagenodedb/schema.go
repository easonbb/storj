//lint:file-ignore * generated file
// AUTOGENERATED BY storj.io/storj/storagenode/storagenodedb/schemagen.go
// DO NOT EDIT

package storagenodedb

import "storj.io/storj/private/dbutil/dbschema"

func Schema() map[string]*dbschema.Schema {
	return map[string]*dbschema.Schema{
		"bandwidth": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name: "bandwidth_usage",
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "action",
							Type:       "INTEGER",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "amount",
							Type:       "BIGINT",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "created_at",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
					},
				},
				&dbschema.Table{
					Name:       "bandwidth_usage_rollups",
					PrimaryKey: []string{"action", "interval_start", "satellite_id"},
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "action",
							Type:       "INTEGER",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "amount",
							Type:       "BIGINT",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "interval_start",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
					},
				},
			},
			Indexes: []*dbschema.Index{
				&dbschema.Index{Name: "idx_bandwidth_usage_created", Table: "bandwidth_usage", Columns: []string{"created_at"}, Unique: false, Partial: ""},
				&dbschema.Index{Name: "idx_bandwidth_usage_satellite", Table: "bandwidth_usage", Columns: []string{"satellite_id"}, Unique: false, Partial: ""},
			},
		},
		"heldamount": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name:       "paystubs",
					PrimaryKey: []string{"period", "satellite_id"},
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "codes",
							Type:       "text",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "comp_at_rest",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "comp_get",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "comp_get_audit",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "comp_get_repair",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "comp_put",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "comp_put_repair",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "created_at",
							Type:       "timestamp",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "disposed",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "held",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "owed",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "paid",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "period",
							Type:       "text",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "bytea",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "surge_percent",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "usage_at_rest",
							Type:       "double precision",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "usage_get",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "usage_get_audit",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "usage_get_repair",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "usage_put",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "usage_put_repair",
							Type:       "bigint",
							IsNullable: false,
						},
					},
				},
			},
		},
		"info": &dbschema.Schema{},
		"notifications": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name:       "notifications",
					PrimaryKey: []string{"id"},
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "created_at",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "message",
							Type:       "TEXT",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "read_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						&dbschema.Column{
							Name:       "sender_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "title",
							Type:       "TEXT",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "type",
							Type:       "INTEGER",
							IsNullable: false,
						},
					},
				},
			},
		},
		"orders": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name: "order_archive_",
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "archived_at",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "order_limit_serialized",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "order_serialized",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "serial_number",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "status",
							Type:       "INTEGER",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "uplink_cert_id",
							Type:       "INTEGER",
							IsNullable: false,
							Reference:  &dbschema.Reference{Table: "certificate", Column: "cert_id", OnDelete: "", OnUpdate: ""},
						},
					},
				},
				&dbschema.Table{
					Name: "unsent_order",
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "order_limit_expiration",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "order_limit_serialized",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "order_serialized",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "serial_number",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "uplink_cert_id",
							Type:       "INTEGER",
							IsNullable: false,
							Reference:  &dbschema.Reference{Table: "certificate", Column: "cert_id", OnDelete: "", OnUpdate: ""},
						},
					},
				},
			},
			Indexes: []*dbschema.Index{
				&dbschema.Index{Name: "idx_order_archived_at", Table: "order_archive_", Columns: []string{"archived_at"}, Unique: false, Partial: ""},
				&dbschema.Index{Name: "idx_orders", Table: "unsent_order", Columns: []string{"satellite_id", "serial_number"}, Unique: false, Partial: ""},
			},
		},
		"piece_expiration": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name:       "piece_expirations",
					PrimaryKey: []string{"piece_id", "satellite_id"},
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "deletion_failed_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						&dbschema.Column{
							Name:       "piece_expiration",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "piece_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "trash",
							Type:       "INTEGER",
							IsNullable: false,
						},
					},
				},
			},
			Indexes: []*dbschema.Index{
				&dbschema.Index{Name: "idx_piece_expirations_deletion_failed_at", Table: "piece_expirations", Columns: []string{"deletion_failed_at"}, Unique: false, Partial: ""},
				&dbschema.Index{Name: "idx_piece_expirations_piece_expiration", Table: "piece_expirations", Columns: []string{"piece_expiration"}, Unique: false, Partial: ""},
				&dbschema.Index{Name: "idx_piece_expirations_trashed", Table: "piece_expirations", Columns: []string{"satellite_id", "trash"}, Unique: false, Partial: "trash = 1"},
			},
		},
		"piece_spaced_used": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name: "piece_space_used",
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "content_size",
							Type:       "INTEGER",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: true,
						},
						&dbschema.Column{
							Name:       "total",
							Type:       "INTEGER",
							IsNullable: false,
						},
					},
				},
			},
			Indexes: []*dbschema.Index{
				&dbschema.Index{Name: "idx_piece_space_used_satellite_id", Table: "piece_space_used", Columns: []string{"satellite_id"}, Unique: false, Partial: ""},
			},
		},
		"pieceinfo": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name: "pieceinfo_",
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "deletion_failed_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						&dbschema.Column{
							Name:       "order_limit",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "piece_creation",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "piece_expiration",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						&dbschema.Column{
							Name:       "piece_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "piece_size",
							Type:       "BIGINT",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "uplink_cert_id",
							Type:       "INTEGER",
							IsNullable: false,
							Reference:  &dbschema.Reference{Table: "certificate", Column: "cert_id", OnDelete: "", OnUpdate: ""},
						},
						&dbschema.Column{
							Name:       "uplink_piece_hash",
							Type:       "BLOB",
							IsNullable: false,
						},
					},
				},
			},
			Indexes: []*dbschema.Index{
				&dbschema.Index{Name: "idx_pieceinfo__expiration", Table: "pieceinfo_", Columns: []string{"piece_expiration"}, Unique: false, Partial: "piece_expiration IS NOT NULL"},
				&dbschema.Index{Name: "pk_pieceinfo_", Table: "pieceinfo_", Columns: []string{"satellite_id", "piece_id"}, Unique: false, Partial: ""},
			},
		},
		"pricing": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name:       "pricing",
					PrimaryKey: []string{"satellite_id"},
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "audit_bandwidth_price",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "disk_space_price",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "egress_bandwidth_price",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "repair_bandwidth_price",
							Type:       "bigint",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
					},
				},
			},
		},
		"reputation": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name:       "reputation",
					PrimaryKey: []string{"satellite_id"},
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "audit_reputation_alpha",
							Type:       "REAL",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "audit_reputation_beta",
							Type:       "REAL",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "audit_reputation_score",
							Type:       "REAL",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "audit_success_count",
							Type:       "INTEGER",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "audit_total_count",
							Type:       "INTEGER",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "disqualified",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						&dbschema.Column{
							Name:       "joined_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "suspended",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						&dbschema.Column{
							Name:       "updated_at",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "uptime_reputation_alpha",
							Type:       "REAL",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "uptime_reputation_beta",
							Type:       "REAL",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "uptime_reputation_score",
							Type:       "REAL",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "uptime_success_count",
							Type:       "INTEGER",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "uptime_total_count",
							Type:       "INTEGER",
							IsNullable: false,
						},
					},
				},
			},
		},
		"satellites": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name:       "satellite_exit_progress",
					PrimaryKey: []string{"satellite_id"},
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "bytes_deleted",
							Type:       "INTEGER",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "completion_receipt",
							Type:       "BLOB",
							IsNullable: true,
						},
						&dbschema.Column{
							Name:       "finished_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						&dbschema.Column{
							Name:       "initiated_at",
							Type:       "TIMESTAMP",
							IsNullable: true,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "starting_disk_usage",
							Type:       "INTEGER",
							IsNullable: false,
						},
					},
				},
				&dbschema.Table{
					Name:       "satellites",
					PrimaryKey: []string{"node_id"},
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "added_at",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "node_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "status",
							Type:       "INTEGER",
							IsNullable: false,
						},
					},
				},
			},
		},
		"storage_usage": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name:       "storage_usage",
					PrimaryKey: []string{"interval_start", "satellite_id"},
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "at_rest_total",
							Type:       "REAL",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "interval_start",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
					},
				},
			},
		},
		"used_serial": &dbschema.Schema{
			Tables: []*dbschema.Table{
				&dbschema.Table{
					Name: "used_serial_",
					Columns: []*dbschema.Column{
						&dbschema.Column{
							Name:       "expiration",
							Type:       "TIMESTAMP",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "satellite_id",
							Type:       "BLOB",
							IsNullable: false,
						},
						&dbschema.Column{
							Name:       "serial_number",
							Type:       "BLOB",
							IsNullable: false,
						},
					},
				},
			},
			Indexes: []*dbschema.Index{
				&dbschema.Index{Name: "idx_used_serial_", Table: "used_serial_", Columns: []string{"expiration"}, Unique: false, Partial: ""},
				&dbschema.Index{Name: "pk_used_serial_", Table: "used_serial_", Columns: []string{"satellite_id", "serial_number"}, Unique: false, Partial: ""},
			},
		},
	}
}
