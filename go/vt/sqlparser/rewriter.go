/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by ASTHelperGen. DO NOT EDIT.

package sqlparser

func (a *application) apply(parent, node SQLNode, replacer replacerFunc) {
	if node == nil || isNilValue(node) {
		return
	}
	saved := a.cursor
	a.cursor.replacer = replacer
	a.cursor.node = node
	a.cursor.parent = parent
	if a.pre != nil && !a.pre(&a.cursor) {
		a.cursor = saved
		return
	}
	switch n := node.(type) {
	case *AddColumns:
		for x, el := range n.Columns {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*AddColumns).Columns[idx] = newNode.(*ColumnDefinition)
				}
			}(x))
		}
		a.apply(node, n.First, func(newNode, parent SQLNode) {
			parent.(*AddColumns).First = newNode.(*ColName)
		})
		a.apply(node, n.After, func(newNode, parent SQLNode) {
			parent.(*AddColumns).After = newNode.(*ColName)
		})
	case *AddConstraintDefinition:
		a.apply(node, n.ConstraintDefinition, func(newNode, parent SQLNode) {
			parent.(*AddConstraintDefinition).ConstraintDefinition = newNode.(*ConstraintDefinition)
		})
	case *AddIndexDefinition:
		a.apply(node, n.IndexDefinition, func(newNode, parent SQLNode) {
			parent.(*AddIndexDefinition).IndexDefinition = newNode.(*IndexDefinition)
		})
	case *AliasedExpr:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*AliasedExpr).Expr = newNode.(Expr)
		})
		a.apply(node, n.As, func(newNode, parent SQLNode) {
			parent.(*AliasedExpr).As = newNode.(ColIdent)
		})
	case *AliasedTableExpr:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*AliasedTableExpr).Expr = newNode.(SimpleTableExpr)
		})
		a.apply(node, n.Partitions, func(newNode, parent SQLNode) {
			parent.(*AliasedTableExpr).Partitions = newNode.(Partitions)
		})
		a.apply(node, n.As, func(newNode, parent SQLNode) {
			parent.(*AliasedTableExpr).As = newNode.(TableIdent)
		})
		a.apply(node, n.Hints, func(newNode, parent SQLNode) {
			parent.(*AliasedTableExpr).Hints = newNode.(*IndexHints)
		})
	case *AlterCharset:
	case *AlterColumn:
		a.apply(node, n.Column, func(newNode, parent SQLNode) {
			parent.(*AlterColumn).Column = newNode.(*ColName)
		})
		a.apply(node, n.DefaultVal, func(newNode, parent SQLNode) {
			parent.(*AlterColumn).DefaultVal = newNode.(Expr)
		})
	case *AlterDatabase:
	case *AlterMigration:
	case *AlterTable:
		a.apply(node, n.Table, func(newNode, parent SQLNode) {
			parent.(*AlterTable).Table = newNode.(TableName)
		})
		for x, el := range n.AlterOptions {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*AlterTable).AlterOptions[idx] = newNode.(AlterOption)
				}
			}(x))
		}
		a.apply(node, n.PartitionSpec, func(newNode, parent SQLNode) {
			parent.(*AlterTable).PartitionSpec = newNode.(*PartitionSpec)
		})
	case *AlterView:
		a.apply(node, n.ViewName, func(newNode, parent SQLNode) {
			parent.(*AlterView).ViewName = newNode.(TableName)
		})
		a.apply(node, n.Columns, func(newNode, parent SQLNode) {
			parent.(*AlterView).Columns = newNode.(Columns)
		})
		a.apply(node, n.Select, func(newNode, parent SQLNode) {
			parent.(*AlterView).Select = newNode.(SelectStatement)
		})
	case *AlterVschema:
		a.apply(node, n.Table, func(newNode, parent SQLNode) {
			parent.(*AlterVschema).Table = newNode.(TableName)
		})
		a.apply(node, n.VindexSpec, func(newNode, parent SQLNode) {
			parent.(*AlterVschema).VindexSpec = newNode.(*VindexSpec)
		})
		for x, el := range n.VindexCols {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*AlterVschema).VindexCols[idx] = newNode.(ColIdent)
				}
			}(x))
		}
		a.apply(node, n.AutoIncSpec, func(newNode, parent SQLNode) {
			parent.(*AlterVschema).AutoIncSpec = newNode.(*AutoIncSpec)
		})
	case *AndExpr:
		a.apply(node, n.Left, func(newNode, parent SQLNode) {
			parent.(*AndExpr).Left = newNode.(Expr)
		})
		a.apply(node, n.Right, func(newNode, parent SQLNode) {
			parent.(*AndExpr).Right = newNode.(Expr)
		})
	case *AutoIncSpec:
		a.apply(node, n.Column, func(newNode, parent SQLNode) {
			parent.(*AutoIncSpec).Column = newNode.(ColIdent)
		})
		a.apply(node, n.Sequence, func(newNode, parent SQLNode) {
			parent.(*AutoIncSpec).Sequence = newNode.(TableName)
		})
	case *Begin:
	case *BinaryExpr:
		a.apply(node, n.Left, func(newNode, parent SQLNode) {
			parent.(*BinaryExpr).Left = newNode.(Expr)
		})
		a.apply(node, n.Right, func(newNode, parent SQLNode) {
			parent.(*BinaryExpr).Right = newNode.(Expr)
		})
	case *CallProc:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*CallProc).Name = newNode.(TableName)
		})
		a.apply(node, n.Params, func(newNode, parent SQLNode) {
			parent.(*CallProc).Params = newNode.(Exprs)
		})
	case *CaseExpr:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*CaseExpr).Expr = newNode.(Expr)
		})
		for x, el := range n.Whens {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*CaseExpr).Whens[idx] = newNode.(*When)
				}
			}(x))
		}
		a.apply(node, n.Else, func(newNode, parent SQLNode) {
			parent.(*CaseExpr).Else = newNode.(Expr)
		})
	case *ChangeColumn:
		a.apply(node, n.OldColumn, func(newNode, parent SQLNode) {
			parent.(*ChangeColumn).OldColumn = newNode.(*ColName)
		})
		a.apply(node, n.NewColDefinition, func(newNode, parent SQLNode) {
			parent.(*ChangeColumn).NewColDefinition = newNode.(*ColumnDefinition)
		})
		a.apply(node, n.First, func(newNode, parent SQLNode) {
			parent.(*ChangeColumn).First = newNode.(*ColName)
		})
		a.apply(node, n.After, func(newNode, parent SQLNode) {
			parent.(*ChangeColumn).After = newNode.(*ColName)
		})
	case *CheckConstraintDefinition:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*CheckConstraintDefinition).Expr = newNode.(Expr)
		})
	case ColIdent:
	case *ColName:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*ColName).Name = newNode.(ColIdent)
		})
		a.apply(node, n.Qualifier, func(newNode, parent SQLNode) {
			parent.(*ColName).Qualifier = newNode.(TableName)
		})
	case *CollateExpr:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*CollateExpr).Expr = newNode.(Expr)
		})
	case *ColumnDefinition:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*ColumnDefinition).Name = newNode.(ColIdent)
		})
	case *ColumnType:
		a.apply(node, n.Length, func(newNode, parent SQLNode) {
			parent.(*ColumnType).Length = newNode.(*Literal)
		})
		a.apply(node, n.Scale, func(newNode, parent SQLNode) {
			parent.(*ColumnType).Scale = newNode.(*Literal)
		})
	case Columns:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(Columns)[idx] = newNode.(ColIdent)
				}
			}(x))
		}
	case Comments:
	case *Commit:
	case *ComparisonExpr:
		a.apply(node, n.Left, func(newNode, parent SQLNode) {
			parent.(*ComparisonExpr).Left = newNode.(Expr)
		})
		a.apply(node, n.Right, func(newNode, parent SQLNode) {
			parent.(*ComparisonExpr).Right = newNode.(Expr)
		})
		a.apply(node, n.Escape, func(newNode, parent SQLNode) {
			parent.(*ComparisonExpr).Escape = newNode.(Expr)
		})
	case *ConstraintDefinition:
		a.apply(node, n.Details, func(newNode, parent SQLNode) {
			parent.(*ConstraintDefinition).Details = newNode.(ConstraintInfo)
		})
	case *ConvertExpr:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*ConvertExpr).Expr = newNode.(Expr)
		})
		a.apply(node, n.Type, func(newNode, parent SQLNode) {
			parent.(*ConvertExpr).Type = newNode.(*ConvertType)
		})
	case *ConvertType:
		a.apply(node, n.Length, func(newNode, parent SQLNode) {
			parent.(*ConvertType).Length = newNode.(*Literal)
		})
		a.apply(node, n.Scale, func(newNode, parent SQLNode) {
			parent.(*ConvertType).Scale = newNode.(*Literal)
		})
	case *ConvertUsingExpr:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*ConvertUsingExpr).Expr = newNode.(Expr)
		})
	case *CreateDatabase:
		a.apply(node, n.Comments, func(newNode, parent SQLNode) {
			parent.(*CreateDatabase).Comments = newNode.(Comments)
		})
	case *CreateTable:
		a.apply(node, n.Table, func(newNode, parent SQLNode) {
			parent.(*CreateTable).Table = newNode.(TableName)
		})
		a.apply(node, n.TableSpec, func(newNode, parent SQLNode) {
			parent.(*CreateTable).TableSpec = newNode.(*TableSpec)
		})
		a.apply(node, n.OptLike, func(newNode, parent SQLNode) {
			parent.(*CreateTable).OptLike = newNode.(*OptLike)
		})
	case *CreateView:
		a.apply(node, n.ViewName, func(newNode, parent SQLNode) {
			parent.(*CreateView).ViewName = newNode.(TableName)
		})
		a.apply(node, n.Columns, func(newNode, parent SQLNode) {
			parent.(*CreateView).Columns = newNode.(Columns)
		})
		a.apply(node, n.Select, func(newNode, parent SQLNode) {
			parent.(*CreateView).Select = newNode.(SelectStatement)
		})
	case *CurTimeFuncExpr:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*CurTimeFuncExpr).Name = newNode.(ColIdent)
		})
		a.apply(node, n.Fsp, func(newNode, parent SQLNode) {
			parent.(*CurTimeFuncExpr).Fsp = newNode.(Expr)
		})
	case *Default:
	case *Delete:
		a.apply(node, n.Comments, func(newNode, parent SQLNode) {
			parent.(*Delete).Comments = newNode.(Comments)
		})
		a.apply(node, n.Targets, func(newNode, parent SQLNode) {
			parent.(*Delete).Targets = newNode.(TableNames)
		})
		a.apply(node, n.TableExprs, func(newNode, parent SQLNode) {
			parent.(*Delete).TableExprs = newNode.(TableExprs)
		})
		a.apply(node, n.Partitions, func(newNode, parent SQLNode) {
			parent.(*Delete).Partitions = newNode.(Partitions)
		})
		a.apply(node, n.Where, func(newNode, parent SQLNode) {
			parent.(*Delete).Where = newNode.(*Where)
		})
		a.apply(node, n.OrderBy, func(newNode, parent SQLNode) {
			parent.(*Delete).OrderBy = newNode.(OrderBy)
		})
		a.apply(node, n.Limit, func(newNode, parent SQLNode) {
			parent.(*Delete).Limit = newNode.(*Limit)
		})
	case *DerivedTable:
		a.apply(node, n.Select, func(newNode, parent SQLNode) {
			parent.(*DerivedTable).Select = newNode.(SelectStatement)
		})
	case *DropColumn:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*DropColumn).Name = newNode.(*ColName)
		})
	case *DropDatabase:
		a.apply(node, n.Comments, func(newNode, parent SQLNode) {
			parent.(*DropDatabase).Comments = newNode.(Comments)
		})
	case *DropKey:
	case *DropTable:
		a.apply(node, n.FromTables, func(newNode, parent SQLNode) {
			parent.(*DropTable).FromTables = newNode.(TableNames)
		})
	case *DropView:
		a.apply(node, n.FromTables, func(newNode, parent SQLNode) {
			parent.(*DropView).FromTables = newNode.(TableNames)
		})
	case *ExistsExpr:
		a.apply(node, n.Subquery, func(newNode, parent SQLNode) {
			parent.(*ExistsExpr).Subquery = newNode.(*Subquery)
		})
	case *ExplainStmt:
		a.apply(node, n.Statement, func(newNode, parent SQLNode) {
			parent.(*ExplainStmt).Statement = newNode.(Statement)
		})
	case *ExplainTab:
		a.apply(node, n.Table, func(newNode, parent SQLNode) {
			parent.(*ExplainTab).Table = newNode.(TableName)
		})
	case Exprs:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(Exprs)[idx] = newNode.(Expr)
				}
			}(x))
		}
	case *Flush:
		a.apply(node, n.TableNames, func(newNode, parent SQLNode) {
			parent.(*Flush).TableNames = newNode.(TableNames)
		})
	case *Force:
	case *ForeignKeyDefinition:
		a.apply(node, n.Source, func(newNode, parent SQLNode) {
			parent.(*ForeignKeyDefinition).Source = newNode.(Columns)
		})
		a.apply(node, n.ReferencedTable, func(newNode, parent SQLNode) {
			parent.(*ForeignKeyDefinition).ReferencedTable = newNode.(TableName)
		})
		a.apply(node, n.ReferencedColumns, func(newNode, parent SQLNode) {
			parent.(*ForeignKeyDefinition).ReferencedColumns = newNode.(Columns)
		})
		a.apply(node, n.OnDelete, func(newNode, parent SQLNode) {
			parent.(*ForeignKeyDefinition).OnDelete = newNode.(ReferenceAction)
		})
		a.apply(node, n.OnUpdate, func(newNode, parent SQLNode) {
			parent.(*ForeignKeyDefinition).OnUpdate = newNode.(ReferenceAction)
		})
	case *FuncExpr:
		a.apply(node, n.Qualifier, func(newNode, parent SQLNode) {
			parent.(*FuncExpr).Qualifier = newNode.(TableIdent)
		})
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*FuncExpr).Name = newNode.(ColIdent)
		})
		a.apply(node, n.Exprs, func(newNode, parent SQLNode) {
			parent.(*FuncExpr).Exprs = newNode.(SelectExprs)
		})
	case GroupBy:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(GroupBy)[idx] = newNode.(Expr)
				}
			}(x))
		}
	case *GroupConcatExpr:
		a.apply(node, n.Exprs, func(newNode, parent SQLNode) {
			parent.(*GroupConcatExpr).Exprs = newNode.(SelectExprs)
		})
		a.apply(node, n.OrderBy, func(newNode, parent SQLNode) {
			parent.(*GroupConcatExpr).OrderBy = newNode.(OrderBy)
		})
		a.apply(node, n.Limit, func(newNode, parent SQLNode) {
			parent.(*GroupConcatExpr).Limit = newNode.(*Limit)
		})
	case *IndexDefinition:
		a.apply(node, n.Info, func(newNode, parent SQLNode) {
			parent.(*IndexDefinition).Info = newNode.(*IndexInfo)
		})
	case *IndexHints:
		for x, el := range n.Indexes {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*IndexHints).Indexes[idx] = newNode.(ColIdent)
				}
			}(x))
		}
	case *IndexInfo:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*IndexInfo).Name = newNode.(ColIdent)
		})
		a.apply(node, n.ConstraintName, func(newNode, parent SQLNode) {
			parent.(*IndexInfo).ConstraintName = newNode.(ColIdent)
		})
	case *Insert:
		a.apply(node, n.Comments, func(newNode, parent SQLNode) {
			parent.(*Insert).Comments = newNode.(Comments)
		})
		a.apply(node, n.Table, func(newNode, parent SQLNode) {
			parent.(*Insert).Table = newNode.(TableName)
		})
		a.apply(node, n.Partitions, func(newNode, parent SQLNode) {
			parent.(*Insert).Partitions = newNode.(Partitions)
		})
		a.apply(node, n.Columns, func(newNode, parent SQLNode) {
			parent.(*Insert).Columns = newNode.(Columns)
		})
		a.apply(node, n.Rows, func(newNode, parent SQLNode) {
			parent.(*Insert).Rows = newNode.(InsertRows)
		})
		a.apply(node, n.OnDup, func(newNode, parent SQLNode) {
			parent.(*Insert).OnDup = newNode.(OnDup)
		})
	case *IntervalExpr:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*IntervalExpr).Expr = newNode.(Expr)
		})
	case *IsExpr:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*IsExpr).Expr = newNode.(Expr)
		})
	case JoinCondition:
		a.apply(node, n.On, replacePanic("JoinCondition On"))
		a.apply(node, n.Using, replacePanic("JoinCondition Using"))
	case *JoinTableExpr:
		a.apply(node, n.LeftExpr, func(newNode, parent SQLNode) {
			parent.(*JoinTableExpr).LeftExpr = newNode.(TableExpr)
		})
		a.apply(node, n.RightExpr, func(newNode, parent SQLNode) {
			parent.(*JoinTableExpr).RightExpr = newNode.(TableExpr)
		})
		a.apply(node, n.Condition, func(newNode, parent SQLNode) {
			parent.(*JoinTableExpr).Condition = newNode.(JoinCondition)
		})
	case *KeyState:
	case *Limit:
		a.apply(node, n.Offset, func(newNode, parent SQLNode) {
			parent.(*Limit).Offset = newNode.(Expr)
		})
		a.apply(node, n.Rowcount, func(newNode, parent SQLNode) {
			parent.(*Limit).Rowcount = newNode.(Expr)
		})
	case ListArg:
	case *Literal:
	case *Load:
	case *LockOption:
	case *LockTables:
	case *MatchExpr:
		a.apply(node, n.Columns, func(newNode, parent SQLNode) {
			parent.(*MatchExpr).Columns = newNode.(SelectExprs)
		})
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*MatchExpr).Expr = newNode.(Expr)
		})
	case *ModifyColumn:
		a.apply(node, n.NewColDefinition, func(newNode, parent SQLNode) {
			parent.(*ModifyColumn).NewColDefinition = newNode.(*ColumnDefinition)
		})
		a.apply(node, n.First, func(newNode, parent SQLNode) {
			parent.(*ModifyColumn).First = newNode.(*ColName)
		})
		a.apply(node, n.After, func(newNode, parent SQLNode) {
			parent.(*ModifyColumn).After = newNode.(*ColName)
		})
	case *Nextval:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*Nextval).Expr = newNode.(Expr)
		})
	case *NotExpr:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*NotExpr).Expr = newNode.(Expr)
		})
	case *NullVal:
	case OnDup:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(OnDup)[idx] = newNode.(*UpdateExpr)
				}
			}(x))
		}
	case *OptLike:
		a.apply(node, n.LikeTable, func(newNode, parent SQLNode) {
			parent.(*OptLike).LikeTable = newNode.(TableName)
		})
	case *OrExpr:
		a.apply(node, n.Left, func(newNode, parent SQLNode) {
			parent.(*OrExpr).Left = newNode.(Expr)
		})
		a.apply(node, n.Right, func(newNode, parent SQLNode) {
			parent.(*OrExpr).Right = newNode.(Expr)
		})
	case *Order:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*Order).Expr = newNode.(Expr)
		})
	case OrderBy:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(OrderBy)[idx] = newNode.(*Order)
				}
			}(x))
		}
	case *OrderByOption:
		a.apply(node, n.Cols, func(newNode, parent SQLNode) {
			parent.(*OrderByOption).Cols = newNode.(Columns)
		})
	case *OtherAdmin:
	case *OtherRead:
	case *ParenSelect:
		a.apply(node, n.Select, func(newNode, parent SQLNode) {
			parent.(*ParenSelect).Select = newNode.(SelectStatement)
		})
	case *ParenTableExpr:
		a.apply(node, n.Exprs, func(newNode, parent SQLNode) {
			parent.(*ParenTableExpr).Exprs = newNode.(TableExprs)
		})
	case *PartitionDefinition:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*PartitionDefinition).Name = newNode.(ColIdent)
		})
		a.apply(node, n.Limit, func(newNode, parent SQLNode) {
			parent.(*PartitionDefinition).Limit = newNode.(Expr)
		})
	case *PartitionSpec:
		a.apply(node, n.Names, func(newNode, parent SQLNode) {
			parent.(*PartitionSpec).Names = newNode.(Partitions)
		})
		a.apply(node, n.Number, func(newNode, parent SQLNode) {
			parent.(*PartitionSpec).Number = newNode.(*Literal)
		})
		a.apply(node, n.TableName, func(newNode, parent SQLNode) {
			parent.(*PartitionSpec).TableName = newNode.(TableName)
		})
		for x, el := range n.Definitions {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*PartitionSpec).Definitions[idx] = newNode.(*PartitionDefinition)
				}
			}(x))
		}
	case Partitions:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(Partitions)[idx] = newNode.(ColIdent)
				}
			}(x))
		}
	case *RangeCond:
		a.apply(node, n.Left, func(newNode, parent SQLNode) {
			parent.(*RangeCond).Left = newNode.(Expr)
		})
		a.apply(node, n.From, func(newNode, parent SQLNode) {
			parent.(*RangeCond).From = newNode.(Expr)
		})
		a.apply(node, n.To, func(newNode, parent SQLNode) {
			parent.(*RangeCond).To = newNode.(Expr)
		})
	case *Release:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*Release).Name = newNode.(ColIdent)
		})
	case *RenameIndex:
	case *RenameTable:
	case *RenameTableName:
		a.apply(node, n.Table, func(newNode, parent SQLNode) {
			parent.(*RenameTableName).Table = newNode.(TableName)
		})
	case *RevertMigration:
	case *Rollback:
	case *SRollback:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*SRollback).Name = newNode.(ColIdent)
		})
	case *Savepoint:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*Savepoint).Name = newNode.(ColIdent)
		})
	case *Select:
		a.apply(node, n.Comments, func(newNode, parent SQLNode) {
			parent.(*Select).Comments = newNode.(Comments)
		})
		a.apply(node, n.SelectExprs, func(newNode, parent SQLNode) {
			parent.(*Select).SelectExprs = newNode.(SelectExprs)
		})
		a.apply(node, n.From, func(newNode, parent SQLNode) {
			parent.(*Select).From = newNode.(TableExprs)
		})
		a.apply(node, n.Where, func(newNode, parent SQLNode) {
			parent.(*Select).Where = newNode.(*Where)
		})
		a.apply(node, n.GroupBy, func(newNode, parent SQLNode) {
			parent.(*Select).GroupBy = newNode.(GroupBy)
		})
		a.apply(node, n.Having, func(newNode, parent SQLNode) {
			parent.(*Select).Having = newNode.(*Where)
		})
		a.apply(node, n.OrderBy, func(newNode, parent SQLNode) {
			parent.(*Select).OrderBy = newNode.(OrderBy)
		})
		a.apply(node, n.Limit, func(newNode, parent SQLNode) {
			parent.(*Select).Limit = newNode.(*Limit)
		})
		a.apply(node, n.Into, func(newNode, parent SQLNode) {
			parent.(*Select).Into = newNode.(*SelectInto)
		})
	case SelectExprs:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(SelectExprs)[idx] = newNode.(SelectExpr)
				}
			}(x))
		}
	case *SelectInto:
	case *Set:
		a.apply(node, n.Comments, func(newNode, parent SQLNode) {
			parent.(*Set).Comments = newNode.(Comments)
		})
		a.apply(node, n.Exprs, func(newNode, parent SQLNode) {
			parent.(*Set).Exprs = newNode.(SetExprs)
		})
	case *SetExpr:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*SetExpr).Name = newNode.(ColIdent)
		})
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*SetExpr).Expr = newNode.(Expr)
		})
	case SetExprs:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(SetExprs)[idx] = newNode.(*SetExpr)
				}
			}(x))
		}
	case *SetTransaction:
		a.apply(node, n.SQLNode, func(newNode, parent SQLNode) {
			parent.(*SetTransaction).SQLNode = newNode.(SQLNode)
		})
		a.apply(node, n.Comments, func(newNode, parent SQLNode) {
			parent.(*SetTransaction).Comments = newNode.(Comments)
		})
		for x, el := range n.Characteristics {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*SetTransaction).Characteristics[idx] = newNode.(Characteristic)
				}
			}(x))
		}
	case *Show:
		a.apply(node, n.Internal, func(newNode, parent SQLNode) {
			parent.(*Show).Internal = newNode.(ShowInternal)
		})
	case *ShowBasic:
		a.apply(node, n.Tbl, func(newNode, parent SQLNode) {
			parent.(*ShowBasic).Tbl = newNode.(TableName)
		})
		a.apply(node, n.Filter, func(newNode, parent SQLNode) {
			parent.(*ShowBasic).Filter = newNode.(*ShowFilter)
		})
	case *ShowCreate:
		a.apply(node, n.Op, func(newNode, parent SQLNode) {
			parent.(*ShowCreate).Op = newNode.(TableName)
		})
	case *ShowFilter:
		a.apply(node, n.Filter, func(newNode, parent SQLNode) {
			parent.(*ShowFilter).Filter = newNode.(Expr)
		})
	case *ShowLegacy:
		a.apply(node, n.OnTable, func(newNode, parent SQLNode) {
			parent.(*ShowLegacy).OnTable = newNode.(TableName)
		})
		a.apply(node, n.Table, func(newNode, parent SQLNode) {
			parent.(*ShowLegacy).Table = newNode.(TableName)
		})
		a.apply(node, n.ShowCollationFilterOpt, func(newNode, parent SQLNode) {
			parent.(*ShowLegacy).ShowCollationFilterOpt = newNode.(Expr)
		})
	case *StarExpr:
		a.apply(node, n.TableName, func(newNode, parent SQLNode) {
			parent.(*StarExpr).TableName = newNode.(TableName)
		})
	case *Stream:
		a.apply(node, n.Comments, func(newNode, parent SQLNode) {
			parent.(*Stream).Comments = newNode.(Comments)
		})
		a.apply(node, n.SelectExpr, func(newNode, parent SQLNode) {
			parent.(*Stream).SelectExpr = newNode.(SelectExpr)
		})
		a.apply(node, n.Table, func(newNode, parent SQLNode) {
			parent.(*Stream).Table = newNode.(TableName)
		})
	case *Subquery:
		a.apply(node, n.Select, func(newNode, parent SQLNode) {
			parent.(*Subquery).Select = newNode.(SelectStatement)
		})
	case *SubstrExpr:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*SubstrExpr).Name = newNode.(*ColName)
		})
		a.apply(node, n.StrVal, func(newNode, parent SQLNode) {
			parent.(*SubstrExpr).StrVal = newNode.(*Literal)
		})
		a.apply(node, n.From, func(newNode, parent SQLNode) {
			parent.(*SubstrExpr).From = newNode.(Expr)
		})
		a.apply(node, n.To, func(newNode, parent SQLNode) {
			parent.(*SubstrExpr).To = newNode.(Expr)
		})
	case TableExprs:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(TableExprs)[idx] = newNode.(TableExpr)
				}
			}(x))
		}
	case TableIdent:
	case TableName:
		a.apply(node, n.Name, replacePanic("TableName Name"))
		a.apply(node, n.Qualifier, replacePanic("TableName Qualifier"))
	case TableNames:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(TableNames)[idx] = newNode.(TableName)
				}
			}(x))
		}
	case TableOptions:
	case *TableSpec:
		for x, el := range n.Columns {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*TableSpec).Columns[idx] = newNode.(*ColumnDefinition)
				}
			}(x))
		}
		for x, el := range n.Indexes {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*TableSpec).Indexes[idx] = newNode.(*IndexDefinition)
				}
			}(x))
		}
		for x, el := range n.Constraints {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*TableSpec).Constraints[idx] = newNode.(*ConstraintDefinition)
				}
			}(x))
		}
		a.apply(node, n.Options, func(newNode, parent SQLNode) {
			parent.(*TableSpec).Options = newNode.(TableOptions)
		})
	case *TablespaceOperation:
	case *TimestampFuncExpr:
		a.apply(node, n.Expr1, func(newNode, parent SQLNode) {
			parent.(*TimestampFuncExpr).Expr1 = newNode.(Expr)
		})
		a.apply(node, n.Expr2, func(newNode, parent SQLNode) {
			parent.(*TimestampFuncExpr).Expr2 = newNode.(Expr)
		})
	case *TruncateTable:
		a.apply(node, n.Table, func(newNode, parent SQLNode) {
			parent.(*TruncateTable).Table = newNode.(TableName)
		})
	case *UnaryExpr:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*UnaryExpr).Expr = newNode.(Expr)
		})
	case *Union:
		a.apply(node, n.FirstStatement, func(newNode, parent SQLNode) {
			parent.(*Union).FirstStatement = newNode.(SelectStatement)
		})
		for x, el := range n.UnionSelects {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*Union).UnionSelects[idx] = newNode.(*UnionSelect)
				}
			}(x))
		}
		a.apply(node, n.OrderBy, func(newNode, parent SQLNode) {
			parent.(*Union).OrderBy = newNode.(OrderBy)
		})
		a.apply(node, n.Limit, func(newNode, parent SQLNode) {
			parent.(*Union).Limit = newNode.(*Limit)
		})
	case *UnionSelect:
		a.apply(node, n.Statement, func(newNode, parent SQLNode) {
			parent.(*UnionSelect).Statement = newNode.(SelectStatement)
		})
	case *UnlockTables:
	case *Update:
		a.apply(node, n.Comments, func(newNode, parent SQLNode) {
			parent.(*Update).Comments = newNode.(Comments)
		})
		a.apply(node, n.TableExprs, func(newNode, parent SQLNode) {
			parent.(*Update).TableExprs = newNode.(TableExprs)
		})
		a.apply(node, n.Exprs, func(newNode, parent SQLNode) {
			parent.(*Update).Exprs = newNode.(UpdateExprs)
		})
		a.apply(node, n.Where, func(newNode, parent SQLNode) {
			parent.(*Update).Where = newNode.(*Where)
		})
		a.apply(node, n.OrderBy, func(newNode, parent SQLNode) {
			parent.(*Update).OrderBy = newNode.(OrderBy)
		})
		a.apply(node, n.Limit, func(newNode, parent SQLNode) {
			parent.(*Update).Limit = newNode.(*Limit)
		})
	case *UpdateExpr:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*UpdateExpr).Name = newNode.(*ColName)
		})
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*UpdateExpr).Expr = newNode.(Expr)
		})
	case UpdateExprs:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(UpdateExprs)[idx] = newNode.(*UpdateExpr)
				}
			}(x))
		}
	case *Use:
		a.apply(node, n.DBName, func(newNode, parent SQLNode) {
			parent.(*Use).DBName = newNode.(TableIdent)
		})
	case *VStream:
		a.apply(node, n.Comments, func(newNode, parent SQLNode) {
			parent.(*VStream).Comments = newNode.(Comments)
		})
		a.apply(node, n.SelectExpr, func(newNode, parent SQLNode) {
			parent.(*VStream).SelectExpr = newNode.(SelectExpr)
		})
		a.apply(node, n.Table, func(newNode, parent SQLNode) {
			parent.(*VStream).Table = newNode.(TableName)
		})
		a.apply(node, n.Where, func(newNode, parent SQLNode) {
			parent.(*VStream).Where = newNode.(*Where)
		})
		a.apply(node, n.Limit, func(newNode, parent SQLNode) {
			parent.(*VStream).Limit = newNode.(*Limit)
		})
	case ValTuple:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(ValTuple)[idx] = newNode.(Expr)
				}
			}(x))
		}
	case *Validation:
	case Values:
		for x, el := range n {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(Values)[idx] = newNode.(ValTuple)
				}
			}(x))
		}
	case *ValuesFuncExpr:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*ValuesFuncExpr).Name = newNode.(*ColName)
		})
	case VindexParam:
		a.apply(node, n.Key, replacePanic("VindexParam Key"))
	case *VindexSpec:
		a.apply(node, n.Name, func(newNode, parent SQLNode) {
			parent.(*VindexSpec).Name = newNode.(ColIdent)
		})
		a.apply(node, n.Type, func(newNode, parent SQLNode) {
			parent.(*VindexSpec).Type = newNode.(ColIdent)
		})
		for x, el := range n.Params {
			a.apply(node, el, func(idx int) func(SQLNode, SQLNode) {
				return func(newNode, container SQLNode) {
					container.(*VindexSpec).Params[idx] = newNode.(VindexParam)
				}
			}(x))
		}
	case *When:
		a.apply(node, n.Cond, func(newNode, parent SQLNode) {
			parent.(*When).Cond = newNode.(Expr)
		})
		a.apply(node, n.Val, func(newNode, parent SQLNode) {
			parent.(*When).Val = newNode.(Expr)
		})
	case *Where:
		a.apply(node, n.Expr, func(newNode, parent SQLNode) {
			parent.(*Where).Expr = newNode.(Expr)
		})
	case *XorExpr:
		a.apply(node, n.Left, func(newNode, parent SQLNode) {
			parent.(*XorExpr).Left = newNode.(Expr)
		})
		a.apply(node, n.Right, func(newNode, parent SQLNode) {
			parent.(*XorExpr).Right = newNode.(Expr)
		})
	}
	if a.post != nil && !a.post(&a.cursor) {
		panic(abort)
	}
	a.cursor = saved
}
