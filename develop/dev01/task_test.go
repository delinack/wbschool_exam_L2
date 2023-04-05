package dev01

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mock_dev01 "main/develop/dev01/mock"
)

func TestTimer_FormatTime(t *testing.T) {
	timer := NewTimer(nil)
	fakeTimeString := "2009-11-10T23:00:00Z"
	fakeTime, err := time.Parse(time.RFC3339Nano, fakeTimeString)
	require.NoError(t, err)

	formattedTime := timer.FormatTime(fakeTime)
	require.Equal(t, fakeTimeString, formattedTime)
}

func TestTimer_GetTime_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	ntpClient := mock_dev01.NewMockINtp(ctrl)

	fakeTimeString := "2009-11-10T23:00:00Z"
	fakeTime, err := time.Parse(time.RFC3339Nano, fakeTimeString)
	require.NoError(t, err)

	ntpClient.EXPECT().GetTime().Return(fakeTime, nil)

	timer := NewTimer(ntpClient)
	currentTime, err := timer.GetTime()
	require.NoError(t, err)
	require.Equal(t, fakeTime, currentTime)
}

func TestTimer_GetTime_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	ntpClient := mock_dev01.NewMockINtp(ctrl)

	ntpClient.EXPECT().GetTime().Return(time.Now(), fmt.Errorf("some error"))

	timer := NewTimer(ntpClient)
	_, err := timer.GetTime()
	require.Error(t, err, "some error")
}
