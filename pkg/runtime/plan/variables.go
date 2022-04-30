/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package plan

import (
	"context"
)

import (
	"github.com/arana-db/arana/pkg/config"
	fieldType "github.com/arana-db/arana/pkg/constants/mysql"
	"github.com/arana-db/arana/pkg/mysql"
	"github.com/arana-db/arana/pkg/proto"
	"github.com/arana-db/arana/pkg/runtime/ast"
	rcontext "github.com/arana-db/arana/pkg/runtime/context"
)

var _ proto.Plan = (*ShowVariablesPlan)(nil)

var (
	_systemSchema = map[config.DataSourceType]string{
		config.DBMySQL: "mysql",
	}
)

type ShowVariablesPlan struct {
	basePlan
	Stmt *ast.ShowVariables
}

func NewShowVariablesPlan(stmt *ast.ShowVariables) *ShowVariablesPlan {
	return &ShowVariablesPlan{Stmt: stmt}
}

func (s *ShowVariablesPlan) Type() proto.PlanType {
	return proto.PlanTypeQuery
}

func (s *ShowVariablesPlan) ExecIn(ctx context.Context, vConn proto.VConn) (proto.Result, error) {
	ret, err := vConn.Query(ctx, "", rcontext.SQL(ctx))
	if err != nil {
		return nil, err
	}

	return &mysql.Result{
		Fields: []proto.Field{mysql.NewField("Variable_name", fieldType.FieldTypeString),
			mysql.NewField("Value", fieldType.FieldTypeString)},
		Rows: ret.GetRows(),
	}, nil
}
